package service

type MealFilter struct {
	Product string
}

func (f *MealFilter) SetProduct(product string) {
	f.Product = product
}
