package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"log"
	"meal-planner/domain"
	"meal-planner/queries"
	"meal-planner/request"
	"net/http"
	"time"
)

var ctx = context.Background()

func main() {
	config, err := NewConfig(
		"localhost",
		5432,
		"fav_food",
		"fav_food",
		"fav_food",
	)

	if err != nil {
		log.Fatalln(err)
	}

	db := NewDatabase(*config)
	cache, _ := bigcache.New(ctx, bigcache.DefaultConfig(10*time.Minute))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Put("/meal-plan", func(w http.ResponseWriter, r *http.Request) {
		database := db.OpenConnection()
		defer database.Close()
		q := queries.New(database)

		req := request.PlanMealRequest{}
		json.NewDecoder(r.Body).Decode(&req)

		parsedDate, _ := time.Parse("2006-01-02", req.Date)
		plan, _ := q.FindPlanByDate(ctx, parsedDate)
		if plan.ID == 0 {
			plan, _ = q.CreatePlan(ctx, parsedDate)
		}

		q.AddMealToPlan(ctx, queries.AddMealToPlanParams{
			PlanID:        plan.ID,
			MealVariantID: req.MealVariantId,
		})
	})

	r.Get("/meals", func(w http.ResponseWriter, r *http.Request) {
		database := db.OpenConnection()
		defer database.Close()
		q := queries.New(database)

		mealModels, err := q.FindAllMeals(ctx)
		var meals []domain.Meal

		for _, mealModel := range mealModels {
			entry, _ := cache.Get("meal" + fmt.Sprintf("%d", mealModel.ID))
			if entry != nil {
				meal := domain.Meal{}
				json.Unmarshal(entry, &meal)
				meals = append(meals, meal)

				continue
			}

			variantModels, err := q.FindMealVariants(ctx, mealModel.ID)
			if err != nil {
				log.Printf("Error while trying to find variants for meal with ID: %d\n", mealModel.ID)
				continue
			}

			var variants []domain.MealVariant
			for _, variantModel := range variantModels {
				macro, err := q.FindMacro(ctx, variantModel.ID)
				if err != nil {
					log.Printf("Error while trying to find macro for meal variant with ID: %d\n", variantModel.ID)
					continue
				}
				ingredientModels, err := q.FindIngredients(ctx, variantModel.ID)
				if err != nil {
					log.Printf("Error while trying to find ingredients for meal variant with ID: %d\n", variantModel.ID)
					continue
				}

				var ingredients []domain.Ingredient
				for _, model := range ingredientModels {
					ingredients = append(ingredients, domain.Ingredient{
						ProductName: model.ProductName.String,
						Amount:      model.Amount,
						Unit:        "g",
						Snack:       model.Snack,
					})
				}

				variants = append(variants, domain.MealVariant{
					ID:        variantModel.ID,
					Kcal:      variantModel.Kcal,
					KcalDaily: variantModel.KcalDaily,
					Macro: domain.Macro{
						Proteins: macro.Proteins,
						Fats:     macro.Fats,
						Carbs:    macro.Carbs,
						Fiber:    macro.Fiber,
					},
					Ingredients: ingredients,
				})
			}

			meal := domain.Meal{
				ID:          mealModel.ID,
				Name:        mealModel.Name,
				Description: mealModel.Description,
				Variants:    variants,
			}
			meals = append(meals, meal)

			mealBytes, err := json.Marshal(meal)
			cache.Set("meal"+fmt.Sprintf("%d", mealModel.ID), mealBytes)
		}

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		render.JSON(w, r, meals)
	})
	http.ListenAndServe(":3000", r)
}
