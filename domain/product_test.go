package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Dilip-Nandakumar/go-shopping-cart/domain"
)

func TestNewProduct(t *testing.T) {
	productName := "Dove Soap"
	product := domain.NewProduct(productName, 39.99)

	assert.Equal(t, productName, product.Name)
	assert.Equal(t, "39.99", product.GetPriceForDisplay())
}
