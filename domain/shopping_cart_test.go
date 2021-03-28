package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Dilip-Nandakumar/go-shopping-cart/domain"
)

func TestNewShoppingCart(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	assert.Equal(t, len(shoppingCart.GetItems()), 0)
}

func TestAddCartItem(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)
	var quantity int64 = 1

	err := shoppingCart.AddCartItem(newProduct, quantity)
	cartItems := shoppingCart.GetItems()

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(cartItems))
	assert.Equal(t, newProduct, cartItems[0].GetProduct())
	assert.Equal(t, "39.99", cartItems[0].GetPriceForDisplay())
	assert.Equal(t, quantity, cartItems[0].GetQuantity())
	assert.Equal(t, "39.99", shoppingCart.GetTotalPriceForDisplay())
}

func TestAddCartItemForMultipleQuantities(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)
	var quantity int64 = 5

	err := shoppingCart.AddCartItem(newProduct, quantity)
	cartItems := shoppingCart.GetItems()

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(cartItems))
	assert.Equal(t, newProduct, cartItems[0].GetProduct())
	assert.Equal(t, "199.95", cartItems[0].GetPriceForDisplay())
	assert.Equal(t, quantity, cartItems[0].GetQuantity())
	assert.Equal(t, "199.95", shoppingCart.GetTotalPriceForDisplay())
}

func TestAddCartItemForMultipleOrdersOfSameProduct(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)

	err1 := shoppingCart.AddCartItem(newProduct, 5)
	err2 := shoppingCart.AddCartItem(newProduct, 3)
	cartItems := shoppingCart.GetItems()

	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
	assert.Equal(t, 1, len(cartItems))
	assert.Equal(t, newProduct, cartItems[0].GetProduct())
	assert.Equal(t, "319.92", cartItems[0].GetPriceForDisplay())
	assert.Equal(t, int64(8), cartItems[0].GetQuantity())
	assert.Equal(t, "319.92", shoppingCart.GetTotalPriceForDisplay())
}

func TestAddCartItemForInvalidInput(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)

	err1 := shoppingCart.AddCartItem(newProduct, 5)
	err2 := shoppingCart.AddCartItem(newProduct, -3)

	expectedErrorMessage := "Error adding quantity to existing cart item of product id: 1: Quantity to add cannot be less than zero"

	assert.Equal(t, nil, err1)
	assert.Equal(t, expectedErrorMessage, err2.Error())
}

func TestTotalPriceForRoundHalfDown(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 0.5649)

	shoppingCart.AddCartItem(newProduct, 1)

	assert.Equal(t, "0.56", shoppingCart.GetTotalPriceForDisplay())
}

func TestTotalPriceForRoundHalfUp(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 0.565)

	shoppingCart.AddCartItem(newProduct, 1)

	assert.Equal(t, "0.57", shoppingCart.GetTotalPriceForDisplay())
}

func TestStringOutput(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)

	shoppingCart.AddCartItem(newProduct, 5)

	assert.Equal(t, shoppingCart.String(), "Name___Quantity___Price\nDove Soaps___5___199.95\nTotal: 199.95")
}
