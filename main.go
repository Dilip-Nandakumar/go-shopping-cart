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

	shoppingCart := domain.NewShoppingCart()

	newProduct := domain.NewProduct(1, "Dove Soaps", 39.99)

	shoppingCart.AddCartItem(newProduct, 5)
	shoppingCart.AddCartItem(newProduct, 3)

	fmt.Println("Shopping Cart:")
	fmt.Println(shoppingCart)
}
