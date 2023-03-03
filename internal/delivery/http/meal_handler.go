package http

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"meal-planner/internal/adapters"
	"meal-planner/internal/service"
	"meal-planner/queries"
)
import nethttp "net/http"

type MealHandler struct {
	mealsService service.MealsService
}

func InitMealHandler(r *chi.Mux, db *sql.DB) {
	q := queries.New(db)

	handler := &MealHandler{
		mealsService: service.NewMealsService(adapters.NewPostgresqlMealRepository(q)),
	}
	r.Get("/meals", handler.listMeals)
}

func (h MealHandler) listMeals(w nethttp.ResponseWriter, r *nethttp.Request) {
	filter := service.MealFilter{}
	product := r.URL.Query().Get("product")
	if product != "" {
		filter.SetProduct(product)
	}

	meals := h.mealsService.ListMeals(r.Context(), filter)
	render.JSON(w, r, meals)
}
