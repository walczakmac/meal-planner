package domain

type Meal struct {
	ID          int16         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Variants    []MealVariant `json:"variants"`
}

func (meal Meal) AddVariant(variant MealVariant) {
	variants := append(meal.Variants, variant)
	meal.Variants = variants
}

type MealVariant struct {
	ID          int16        `json:"id"`
	Kcal        float32      `json:"kcal"`
	KcalDaily   int16        `json:"kcal_daily"`
	Macro       Macro        `json:"macro"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ProductName string `json:"product_name"`
	Amount      int16  `json:"amount"`
	Unit        string `json:"unit"`
	Snack       bool   `json:"snack"`
}

type Macro struct {
	Proteins float32 `json:"proteins"`
	Fats     float32 `json:"fats"`
	Carbs    float32 `json:"carbs"`
	Fiber    float32 `json:"fiber"`
}
