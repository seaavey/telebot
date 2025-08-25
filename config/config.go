package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BotToken string
	Owners   []int64
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	BotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set. Please set it in your .env file or as an environment variable.")
	}
	
	Owners = []int64{5716661796}
}