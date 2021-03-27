package main

import (
	"github.com/Dilip-Nandakumar/go-shopping-cart/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()

	log.Info("Main method is invoked")
}
