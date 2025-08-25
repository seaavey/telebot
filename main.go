package main

import (
	"log"
	"telebot/config"
	"telebot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	messageHandler := handlers.NewMessageHandler()

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go messageHandler.HandleUpdate(update, bot)
	}
}