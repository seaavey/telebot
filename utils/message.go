package utils

import (
	"telebot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageType string

const (
	TextMessage     MessageType = "text"
	PhotoMessage    MessageType = "photo"
	VideoMessage    MessageType = "video"
	DocumentMessage MessageType = "document"
	AudioMessage    MessageType = "audio"
	VoiceMessage    MessageType = "voice"
	StickerMessage  MessageType = "sticker"
)

type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
}

type Chat struct {
	ID        int64
	Type      string
	Title     string
	UserName  string
	FirstName string
	LastName  string
}

type Message struct {
	ID        int
	From      *User
	Chat      *Chat
	Type      MessageType
	Text      string
	Date      int
	client    *tgbotapi.BotAPI
}

func (msg *Message) Owner() bool {
    for _, id := range config.Owners {
        if msg.From != nil && msg.From.ID == id {
            return true
        }
    }
    return false
}
func (msg *Message) SetBotAPI(bot *tgbotapi.BotAPI) {
	msg.client = bot
}

func (msg *Message) SendText(text string) error {
	if msg.client == nil {
		return nil
	}
	
	message := tgbotapi.NewMessage(msg.Chat.ID, text)
	_, err := msg.client.Send(message)
	return err
}

func (msg *Message) SendImage(imageURL string, caption ...string) error {
	if msg.client == nil {
		return nil
	}
	
	// Use URL directly instead of FileID for external images
	message := tgbotapi.NewPhoto(msg.Chat.ID, tgbotapi.FileURL(imageURL))

	if len(caption) > 0 {
		message.Caption = caption[0]
	}
	
	_, err := msg.client.Send(message)
	return err
}

func (msg *Message) SendVideo(videoURL string, caption ...string) error {
	if msg.client == nil {
		return nil
	}

	// Use URL directly instead of FileID for external videos
	message := tgbotapi.NewVideo(msg.Chat.ID, tgbotapi.FileURL(videoURL))
	
	if len(caption) > 0 {
		message.Caption = caption[0]
	}

	_, err := msg.client.Send(message)
	return err
}


func (msg *Message) SendTextWithButton(text string, buttons ...tgbotapi.InlineKeyboardButton) error {
	if msg.client == nil {
		return nil
	}
	
	message := tgbotapi.NewMessage(msg.Chat.ID, text)
	
	if len(buttons) > 0 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
		message.ReplyMarkup = &keyboard
	}
	
	_, err := msg.client.Send(message)
	return err
}