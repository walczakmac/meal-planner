package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"meal-planner/internal/delivery/http"
	"meal-planner/pkg/database"
	"meal-planner/pkg/server"
)

func Run() {
	db := database.OpenConnection(
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
	)
	defer db.Close()

	//cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	r := chi.NewRouter()
	s := server.New(r)

	http.InitMealHandler(r, db)

	err := s.Start()
	if err != nil {
		log.Fatalln(err)
	}

	//r.Get("/", api.FindMealsByProductHandler)
	//
	//r.Get("/meals/by-product", func(w http.ResponseWriter, r *http.Request) {
	//	database := db.OpenConnection()
	//	defer database.Close()
	//	q := queries.New(database)
	//
	//	productName := fmt.Sprintf("%%%s%%", r.URL.Query().Get("product-name"))
	//	log.Println(productName)
	//
	//	meals, err := q.FindMealWithByProductName(ctx, productName)
	//	if err != nil {
	//		log.Println(err)
	//		log.Printf("Couldn't find meals with product %s", productName)
	//	}
	//
	//	render.JSON(w, r, meals)
	//})
	//
	//r.Put("/meal-plan", func(w http.ResponseWriter, r *http.Request) {
	//	database := db.OpenConnection()
	//	defer database.Close()
	//	q := queries.New(database)
	//
	//	req := request.PlanMealRequest{}
	//	json.NewDecoder(r.Body).Decode(&req)
	//
	//	parsedDate, _ := time.Parse("2006-01-02", req.Date)
	//	plan, _ := q.FindPlanByDate(ctx, parsedDate)
	//	if plan.ID == 0 {
	//		plan, _ = q.CreatePlan(ctx, parsedDate)
	//	}
	//
	//	q.AddMealToPlan(ctx, queries.AddMealToPlanParams{
	//		PlanID:        plan.ID,
	//		MealVariantID: req.MealVariantId,
	//	})
	//})
	//
	//r.Get("/meals", func(w http.ResponseWriter, r *http.Request) {
	//	database := db.OpenConnection()
	//	defer database.Close()
	//	q := queries.New(database)
	//
	//	mealModels, err := q.FindAllMeals(ctx)
	//	var meals []domain.Meal
	//
	//	for _, mealModel := range mealModels {
	//		entry, _ := cache.Get("meal" + fmt.Sprintf("%d", mealModel.ID))
	//		if entry != nil {
	//			meal := domain.Meal{}
	//			json.Unmarshal(entry, &meal)
	//			meals = append(meals, meal)
	//
	//			continue
	//		}
	//
	//		variantModels, err := q.FindMealVariants(ctx, mealModel.ID)
	//		if err != nil {
	//			log.Printf("Error while trying to find variants for meal with ID: %d\n", mealModel.ID)
	//			continue
	//		}
	//
	//		var variants []domain.MealVariant
	//		for _, variantModel := range variantModels {
	//			macro, err := q.FindMacro(ctx, variantModel.ID)
	//			if err != nil {
	//				log.Printf("Error while trying to find macro for meal variant with ID: %d\n", variantModel.ID)
	//				continue
	//			}
	//			ingredientModels, err := q.FindIngredients(ctx, variantModel.ID)
	//			if err != nil {
	//				log.Printf("Error while trying to find ingredients for meal variant with ID: %d\n", variantModel.ID)
	//				continue
	//			}
	//
	//			var ingredients []domain.Ingredient
	//			for _, model := range ingredientModels {
	//				ingredients = append(ingredients, domain.Ingredient{
	//					ProductName: model.ProductName.String,
	//					Amount:      model.Amount,
	//					Unit:        "g",
	//					Snack:       model.Snack,
	//				})
	//			}
	//
	//			variants = append(variants, domain.MealVariant{
	//				ID:        variantModel.ID,
	//				Kcal:      variantModel.Kcal,
	//				KcalDaily: variantModel.KcalDaily,
	//				Macro: domain.Macro{
	//					Proteins: macro.Proteins,
	//					Fats:     macro.Fats,
	//					Carbs:    macro.Carbs,
	//					Fiber:    macro.Fiber,
	//				},
	//				Ingredients: ingredients,
	//			})
	//		}
	//
	//		meal := domain.Meal{
	//			ID:          mealModel.ID,
	//			Name:        mealModel.Name,
	//			Description: mealModel.Description,
	//			Variants:    variants,
	//		}
	//		meals = append(meals, meal)
	//
	//		mealBytes, err := json.Marshal(meal)
	//		cache.Set("meal"+fmt.Sprintf("%d", mealModel.ID), mealBytes)
	//	}
	//
	//	if err != nil {
	//		w.Write([]byte(err.Error()))
	//	}
	//
	//	render.JSON(w, r, meals)
	//})

	//http.ListenAndServe(":3000", r)
}
