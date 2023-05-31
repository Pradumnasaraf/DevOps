package main

import (
	"github.com/Pradumnasaraf/discord-bot/bot"
	"github.com/Pradumnasaraf/discord-bot/config"
)

func main() {

	config.LoadConfig()

	bot.Start()

	<-make(chan struct{})	// Block main() from exiting
	return
}
