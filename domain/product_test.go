package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Dilip-Nandakumar/go-shopping-cart/domain"
)

func TestNewProduct(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(1, productName, 39.99)

	assert.Equal(t, 1, product.GetID())
	assert.Equal(t, productName, product.GetName())
	assert.Equal(t, "39.99", product.GetPriceForDisplay())
}
