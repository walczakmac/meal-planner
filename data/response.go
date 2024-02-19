package data

type Ingredient struct {
	ProductName string `json:"product_name"`
	Amount      int16  `json:"amount"`
	Unit        string `json:"unit"`
}

type Macro struct {
	Proteins float64 `json:"proteins"`
	Fats     float64 `json:"fats"`
	Carbs    float64 `json:"carbs"`
	Fiber    float64 `json:"fiber"`
}

type MealResponse struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	MealVariantId int16        `json:"meal_variant_id"`
	Kcal          float64      `json:"kcal"`
	Macro         Macro        `json:"macro"`
	Ingredients   []Ingredient `json:"ingredients"`
	Snacks        []Ingredient `json:"snacks"`
}

type MealListRow struct {
	Name     string            `json:"name"`
	Variants []MealListVariant `json:"variants"`
}
type MealListVariant struct {
	MealVariantId int16   `json:"meal_variant_id"`
	Kcal          float32 `json:"kcal"`
}
