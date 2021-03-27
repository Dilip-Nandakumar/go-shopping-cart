package domain

import (
	"github.com/shopspring/decimal"
)

type product struct {
	Name  string
	price decimal.Decimal
}

func NewProduct(name string, price float64) *product {
	return &product{name, decimal.NewFromFloat(price)}
}

func (product *product) GetPrice() decimal.Decimal {
	return product.price.Round(2)
}

func (product *product) GetPriceForDisplay() string {
	return product.GetPrice().String()
}
