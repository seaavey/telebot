package main

import (
	"telebot/config"
	"telebot/handlers"
	"telebot/logger"
	"telebot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		logger.Fatalf("Failed to create bot: %s", err)
	}

	bot.Debug = utils.GetEnvAsBool("BOT_DEBUG", false)

	if bot.Debug {
		logger.Info("Debug mode is ENABLED")
	} else {
		logger.Info("Debug mode is DISABLED")
	}

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	messageHandler := handlers.NewMessageHandler()

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go messageHandler.HandleUpdate(update, bot)
	}
}