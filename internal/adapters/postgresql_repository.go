package adapters

import (
	"context"
	"meal-planner/internal/domain"
	"meal-planner/queries"
)

type PostgresqlMealRepository struct {
	queries *queries.Queries
}

func NewPostgresqlMealRepository(q *queries.Queries) PostgresqlMealRepository {
	return PostgresqlMealRepository{queries: q}
}

func (r PostgresqlMealRepository) FindMealsByProduct(ctx context.Context, productName string) (meals []domain.Meal) {
	rows, err := r.queries.FindMealByProductName(ctx, productName)
	if err != nil {
		return
	}

	for _, meal := range rows {
		meals = append(meals, domain.Meal{
			ID:          meal.ID,
			Name:        meal.Name,
			Description: meal.Description,
			Variants:    nil,
		})
	}

	return
}
