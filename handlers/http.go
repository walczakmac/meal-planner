package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"meal-planner/data"
	"meal-planner/queries"
	"net/http"
	"strings"
	"time"
)

func PlanAMeal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := queries.New(db)

		req := data.PlanMealRequest{}
		json.NewDecoder(r.Body).Decode(&req)

		date := chi.URLParam(r, "date")
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		plan, _ := q.FindPlanByDate(r.Context(), parsedDate)
		if plan.ID == 0 {
			plan, err = q.CreatePlan(r.Context(), parsedDate)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}

		err = q.AddMealToPlan(r.Context(), queries.AddMealToPlanParams{
			PlanID:        plan.ID,
			MealVariantID: req.MealVariantId,
		})
		if err != nil {
			slog.Error(
				"Couldn't create meal",
				"planId", plan.ID,
				"variantId", req.MealVariantId,
				"date", parsedDate,
			)
			return
		}

		http.StatusText(201)
	}
}

func GetMealPlanByDate(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp []data.MealResponse
		date := chi.URLParam(r, "date")
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		q := queries.New(db)
		mealVariants, err := q.FindMealVariantsByDate(r.Context(), parsedDate)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		for _, mealVariant := range mealVariants {
			ingredients, err := q.FindIngredientsByMealVariantId(r.Context(), mealVariant.MealVariantID.Int16)
			if err != nil {
				slog.Error(
					"Error happened when looking up ingredients for meal variant.",
					"mealVariantId",
					mealVariant.MealVariantID,
				)
				continue
			}

			meal := data.MealResponse{
				Name:          mealVariant.Name.String,
				Description:   strings.ReplaceAll(mealVariant.Description.String, "\n", "<br>"),
				MealVariantId: mealVariant.MealVariantID.Int16,
				Kcal:          mealVariant.Kcal.Float64,
				Macro: data.Macro{
					Proteins: mealVariant.Proteins.Float64,
					Fats:     mealVariant.Fats.Float64,
					Carbs:    mealVariant.Carbs.Float64,
					Fiber:    mealVariant.Fiber.Float64,
				},
				Ingredients: nil,
				Snacks:      nil,
			}

			for _, ingredient := range ingredients {
				ing := data.Ingredient{
					ProductName: ingredient.Name.String,
					Amount:      ingredient.Amount,
					Unit:        ingredient.Unit,
				}

				if ingredient.Snack {
					meal.Snacks = append(meal.Snacks, ing)
				} else {
					meal.Ingredients = append(meal.Ingredients, ing)
				}
			}

			resp = append(resp, meal)
		}

		if resp == nil {
			resp = []data.MealResponse{}
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			slog.Error("Couldn't encode response to json", r.URL.String())
			return
		}
	}
}
