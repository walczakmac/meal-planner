package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"log"
	"meal-planner/database"
	"meal-planner/handlers"
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

	//cache, _ := bigcache.NewServer(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	router := chi.NewRouter()
	s := NewServer(router)

	//router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	//	err := ui.Index().Render(context.Background(), writer)
	//	if err != nil {
	//		return
	//	}
	//})
	router.Get("/", handlers.HandleListMeals(db))
	router.Get("/meal/{id}", handlers.HandleShowMeal(db))
	router.Post("/plan/{date}", handlers.PlanAMeal(db))
	router.Get("/plan/{date}", handlers.GetMealPlanByDate(db))

	err := s.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
