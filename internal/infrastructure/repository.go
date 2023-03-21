package infrastructure

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
			Variants:    r.findVariants(ctx, meal.ID),
		})
	}

	return
}

func (r PostgresqlMealRepository) FindAllMeals(ctx context.Context) (meals []domain.Meal) {
	rows, err := r.queries.FindAllMeals(ctx)
	if err != nil {
		return
	}

	for _, meal := range rows {
		meals = append(meals, domain.Meal{
			ID:          meal.ID,
			Name:        meal.Name,
			Description: meal.Description,
			Variants:    r.findVariants(ctx, meal.ID),
		})
	}

	return
}

func (r PostgresqlMealRepository) findVariants(ctx context.Context, mealId int16) []domain.MealVariant {
	variants, err := r.queries.FindMealVariants(ctx, mealId)
	var variantModels []domain.MealVariant
	if err != nil {
		return variantModels
	}

	for _, variant := range variants {
		variantModels = append(variantModels, domain.MealVariant{
			ID:        variant.ID,
			Kcal:      variant.Kcal,
			KcalDaily: variant.KcalDaily,
			Macro: domain.Macro{
				Proteins: variant.Proteins,
				Fats:     variant.Fats,
				Carbs:    variant.Carbs,
				Fiber:    variant.Fiber,
			},
			Ingredients: r.findIngredients(ctx, variant.ID),
		})
	}

	return variantModels
}

func (r PostgresqlMealRepository) findIngredients(ctx context.Context, mealVariantId int16) []domain.Ingredient {
	ingredients, err := r.queries.FindIngredients(ctx, mealVariantId)
	var ingredientModels []domain.Ingredient
	if err != nil {
		return ingredientModels
	}

	for _, ingredient := range ingredients {
		ingredientModels = append(ingredientModels, domain.Ingredient{
			ProductName: ingredient.Name,
			Amount:      ingredient.Amount,
			Unit:        ingredient.Unit,
			Snack:       ingredient.Snack,
		})
	}

	return ingredientModels
}
