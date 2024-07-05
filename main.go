package main

import (
	"fmt"
	"github.com/trynax/inj-price-checker/bot"
	"github.com/trynax/inj-price-checker/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
