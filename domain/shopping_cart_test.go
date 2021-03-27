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

	newProduct := domain.NewProduct("Dove Soaps", 39.99)
	var quantity int64 = 1

	shoppingCart.AddCartItem(newProduct, quantity)
	cartItems := shoppingCart.GetItems()

	assert.Equal(t, 1, len(cartItems))
	assert.Equal(t, newProduct, cartItems[0].GetProduct())
	assert.Equal(t, "39.99", cartItems[0].GetPriceForDisplay())
	assert.Equal(t, quantity, cartItems[0].GetQuantity())
	assert.Equal(t, "39.99", shoppingCart.GetTotalPriceForDisplay())
}

func TestAddCartItemForMultipleQuantities(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct("Dove Soaps", 39.99)
	var quantity int64 = 5

	shoppingCart.AddCartItem(newProduct, quantity)
	cartItems := shoppingCart.GetItems()

	assert.Equal(t, 1, len(cartItems))
	assert.Equal(t, newProduct, cartItems[0].GetProduct())
	assert.Equal(t, "199.95", cartItems[0].GetPriceForDisplay())
	assert.Equal(t, quantity, cartItems[0].GetQuantity())
	assert.Equal(t, "199.95", shoppingCart.GetTotalPriceForDisplay())
}

func TestTotalPriceForRoundHalfDown(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct("Dove Soaps", 0.5649)

	shoppingCart.AddCartItem(newProduct, 1)

	assert.Equal(t, "0.56", shoppingCart.GetTotalPriceForDisplay())
}

func TestTotalPriceForRoundHalfUp(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct("Dove Soaps", 0.565)

	shoppingCart.AddCartItem(newProduct, 1)

	assert.Equal(t, "0.57", shoppingCart.GetTotalPriceForDisplay())
}

func TestStringOutput(t *testing.T) {
	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct("Dove Soaps", 39.99)

	shoppingCart.AddCartItem(newProduct, 5)

	assert.Equal(t, shoppingCart.String(), "Name___Quantity___Price\nDove Soaps___5___199.95\nTotal: 199.95")
}
