package main

import (
	"log"
	"telebot/config"
	"telebot/handlers"
	"telebot/utils" // Tambahkan import untuk utils

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	bot.Debug = utils.GetEnvAsBool("BOT_DEBUG", false)

	if bot.Debug {
		log.Println("Debug mode is ENABLED")
	} else {
		log.Println("Debug mode is DISABLED")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	messageHandler := handlers.NewMessageHandler()

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go messageHandler.HandleUpdate(update, bot)
	}
}