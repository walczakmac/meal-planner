package domain

import "context"

type MealRepository interface {
	FindMealsByProduct(ctx context.Context, productName string) (meals []Meal)
	FindAllMeals(ctx context.Context) (meals []Meal)
}
