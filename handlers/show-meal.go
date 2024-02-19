package handlers

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"meal-planner/queries"
	"meal-planner/ui"
	"net/http"
	"strconv"
)

func HandleShowMeal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := queries.New(db)
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		a, err := q.FindOneMealById(r.Context(), int16(id))
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		err = ui.Meal(a.Name, a.Description).Render(context.Background(), w)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}
