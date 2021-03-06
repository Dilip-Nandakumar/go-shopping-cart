package domain

import (
	"errors"

	"github.com/shopspring/decimal"
)

type cartItem struct {
	product  *product
	quantity int64
}

func NewCartItem(product *product, quantity int64) *cartItem {
	return &cartItem{product, quantity}
}

func (cartItem *cartItem) GetPrice() decimal.Decimal {
	return cartItem.product.GetPrice().Mul(decimal.NewFromInt(cartItem.quantity))
}

func (cartItem *cartItem) GetPriceForDisplay() string {
	return cartItem.GetPrice().String()
}

func (cartItem *cartItem) GetProduct() *product {
	return cartItem.product
}

func (cartItem *cartItem) GetQuantity() int64 {
	return cartItem.quantity
}

func (cartItem *cartItem) AddQuantity(quantity int64) error {
	if quantity < 0 {
		return errors.New("Quantity to add cannot be less than zero")
	}

	cartItem.quantity += quantity
	return nil
}
