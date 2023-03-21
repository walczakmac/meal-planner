package http

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"meal-planner/internal/infrastructure"
	"meal-planner/internal/service"
	"meal-planner/queries"
)
import "net/http"

type MealHandler struct {
	mealsService service.MealsService
}

func InitMealHandler(r *chi.Mux, db *sql.DB) {
	handler := &MealHandler{
		mealsService: service.NewMealsService(
			infrastructure.NewPostgresqlMealRepository(queries.New(db)),
		),
	}
	r.Get("/meals", handler.listMeals)
}

func (h MealHandler) listMeals(w http.ResponseWriter, r *http.Request) {
	filter := service.MealFilter{}
	product := r.URL.Query().Get("product")
	if product != "" {
		filter.SetProduct(product)
	}

	meals := h.mealsService.ListMeals(r.Context(), filter)
	render.JSON(w, r, meals)
}
