package main

import (
	"fmt"
	"test/testBot/bot"
	"test/testBot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Init()

	<-make(chan struct{})
	return
}
