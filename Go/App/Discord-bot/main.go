package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Pradumnasaraf/discord-bot/bot"
	"github.com/Pradumnasaraf/discord-bot/config"
)

func main() {

	config.LoadConfig()

	bot.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
