package domain

import (
	"fmt"
	"github.com/pkg/errors"
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

func (shoppingCart *shoppingCart) AddCartItem(product *product, quantity int64) error {
	if cartItem, ok := shoppingCart.getCartItem(product.GetID()); ok {
		err := cartItem.AddQuantity(quantity)
		return errors.Wrapf(
			err, "Error adding quantity to existing cart item of product id: %d",
			cartItem.product.GetID(),
		)
	} else {
		cartItem := NewCartItem(product, quantity)
		shoppingCart.items = append(shoppingCart.items, cartItem)
	}

	log.Infof("Added item %s of quantity %d to shopping cart", product.GetName(), quantity)
	return nil
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
		output.WriteString(fmt.Sprintf("%s___%d___%s\n", item.product.GetName(), item.quantity, item.GetPriceForDisplay()))
	}

	output.WriteString(fmt.Sprintf("Total: %s", shoppingCart.GetTotalPriceForDisplay()))

	return output.String()
}

func  (shoppingCart *shoppingCart) getCartItem(id int) (*cartItem, bool) {
	for _, item := range shoppingCart.GetItems() {
		if item.product.GetID() == id {
			return item, true
		}
	}

	return nil, false
}
