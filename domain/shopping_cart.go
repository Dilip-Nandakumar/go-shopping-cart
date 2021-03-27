package domain

import (
	"fmt"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"strings"
)

type shoppingCart struct {
	items      []*cartItem
}

func NewShoppingCart() *shoppingCart {
	return &shoppingCart{[]*cartItem{}}
}

func (shoppingCart *shoppingCart) AddCartItem(product *product, quantity int64) {
	cartItem := NewCartItem(product, quantity)
	shoppingCart.items = append(shoppingCart.items, cartItem)
	log.Infof("Added item %s of quantity %d to shopping cart", product.Name, quantity)
}

func (shoppingCart *shoppingCart) GetItems() []*cartItem {
	return shoppingCart.items
}

func (shoppingCart *shoppingCart) GetTotalPrice() decimal.Decimal {
	totalPrice := decimal.NewFromInt(0)

	for _, item := range shoppingCart.items {
		totalPrice = totalPrice.Add(item.GetPrice())
	}

	return totalPrice.Round(2)
}

func (shoppingCart *shoppingCart) GetTotalPriceForDisplay() string {
	return shoppingCart.GetTotalPrice().String()
}

func (shoppingCart *shoppingCart) String() string {
	var output strings.Builder
	output.WriteString("Name___Quantity___Price\n")

	for _, item := range shoppingCart.items {
		output.WriteString(fmt.Sprintf("%s___%d___%s\n", item.product.Name, item.quantity, item.GetPriceForDisplay()))
	}

	output.WriteString(fmt.Sprintf("Total: %s", shoppingCart.GetTotalPriceForDisplay()))

	return output.String()
}
