package main

import (
	"github.com/Dilip-Nandakumar/go-shopping-cart/domain"
	"github.com/Dilip-Nandakumar/go-shopping-cart/utils"

	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()

	log.Info("Main method is invoked")

	shoppingCart := domain.NewShoppingCart(12.5)

	product1 := domain.NewProduct(1, "Dove Soaps", 39.99)
	product2 := domain.NewProduct(2, "Axe Deos", 99.99)

	shoppingCart.AddCartItem(product1, 2)
	shoppingCart.AddCartItem(product2, 2)

	fmt.Println("Shopping Cart:")
	fmt.Println(shoppingCart)
}
