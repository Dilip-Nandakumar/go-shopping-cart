package domain

import (
	"github.com/shopspring/decimal"
)

type product struct {
	id    int
	name  string
	price decimal.Decimal
}

func NewProduct(id int, name string, price float64) *product {
	return &product{id, name, decimal.NewFromFloat(price)}
}

func (product *product) GetName() string {
	return product.name
}

func (product *product) GetID() int {
	return product.id
}

func (product *product) GetPrice() decimal.Decimal {
	return product.price.Round(2)
}

func (product *product) GetPriceForDisplay() string {
	return product.GetPrice().String()
}
