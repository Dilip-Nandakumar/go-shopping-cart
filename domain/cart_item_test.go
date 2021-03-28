package domain_test

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Dilip-Nandakumar/go-shopping-cart/domain"
)

func TestNewCartItem(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)
	var quantity int64 = 1

	cartItem := domain.NewCartItem(product, 1)

	assert.Equal(t, product, cartItem.GetProduct())
	assert.Equal(t, quantity, cartItem.GetQuantity())
}

func TestPrice(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)

	cartItem := domain.NewCartItem(product, 1)

	assert.Equal(t, decimal.NewFromFloat(39.99), cartItem.GetPrice())
	assert.Equal(t, "39.99", cartItem.GetPriceForDisplay())
}

func TestPriceForMultipeQuantities(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)

	cartItem := domain.NewCartItem(product, 5)

	assert.Equal(t, decimal.NewFromFloat(199.95), cartItem.GetPrice())
	assert.Equal(t, "199.95", cartItem.GetPriceForDisplay())
}

func TestAddQuantity(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)

	cartItem := domain.NewCartItem(product, 5)
	err := cartItem.AddQuantity(3)

	assert.Equal(t, int64(8), cartItem.GetQuantity())
	assert.Equal(t, nil, err)
}

func TestAddQuantityForInvalidInput(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)

	cartItem := domain.NewCartItem(product, 5)
	err := cartItem.AddQuantity(-3)

	assert.Equal(t, int64(5), cartItem.GetQuantity())
	assert.Equal(t, "Quantity to add cannot be less than zero", err.Error())
}
