package handlers

import (
	"context"
	"database/sql"
	"log/slog"
	"meal-planner/data"
	"meal-planner/queries"
	"meal-planner/ui"
	"net/http"
)

func HandleListMeals(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := queries.New(db)
		resp := make(map[int]data.MealListRow)
		mealVariants, err := q.FindAllMealsVariantsWithKcal(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			slog.Error("Couldn't find meal variants with kcal", "error", err)
			return
		}

		for i, variant := range mealVariants {
			row, exists := resp[i]
			if !exists {
				row = data.MealListRow{
					Name:     variant.Name,
					Variants: nil,
				}
			}

			row.Variants = append(row.Variants, data.MealListVariant{
				MealVariantId: variant.MealVariantID,
				Kcal:          variant.Kcal,
			})

			resp[i] = row
		}

		v := make([]data.MealListRow, 0, len(resp))
		for _, value := range resp {
			v = append(v, value)
		}

		err = ui.Index(v).Render(context.Background(), w)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		//err = json.NewEncoder(w).Encode(v)
		//if err != nil {
		//	slog.Error("Couldn't encode response to json", r.URL.String())
		//	return
		//}
	}
}
