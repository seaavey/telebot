package config

import (
	"os"
	"strconv"
	"strings"

	"telebot/logger"

	"github.com/joho/godotenv"
)

var (
	BotToken string
	Owners   []int64
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Warning: Error loading .env file")
	}

	BotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if BotToken == "" {
		logger.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set. Please set it in your .env file or as an environment variable.")
	}
	
	// Load owners from environment variable
	ownersStr := os.Getenv("BOT_OWNERS")
	if ownersStr != "" {
		ownerIDs := strings.Split(ownersStr, ",")
		for _, idStr := range ownerIDs {
			if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
				Owners = append(Owners, id)
			}
		}
	} else {
		// Default owner
		Owners = []int64{5716661796}
	}
}