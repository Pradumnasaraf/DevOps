package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	Token = os.Getenv("DISCORD_BOT_TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")
}
