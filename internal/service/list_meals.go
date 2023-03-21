package service

import (
	"context"
	"meal-planner/internal/domain"
)

type MealsService struct {
	meals domain.MealRepository
}

func NewMealsService(repository domain.MealRepository) MealsService {
	return MealsService{
		meals: repository,
	}
}

func (u MealsService) ListMeals(ctx context.Context, filter MealFilter) []domain.Meal {
	if filter.Product != "" {
		return u.meals.FindMealsByProduct(ctx, filter.Product)
	}

	return u.meals.FindAllMeals(ctx)
}
