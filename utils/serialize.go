package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SerializeMessage(message *tgbotapi.Message, botAPI *tgbotapi.BotAPI) *Message {
	if message == nil {
		return nil
	}

	msg := &Message{
		ID:      message.MessageID,
		Text: 	Coalesce(message.Text, message.Caption),
		Date:    message.Date,
		Type:    getMessageType(message),
	}

	msg.SetBotAPI(botAPI)

	if message.From != nil {
		msg.From = &User{
			ID:        message.From.ID,
			UserName:  message.From.UserName,
			FirstName: message.From.FirstName,
			LastName:  message.From.LastName,
		}
	}

	if message.Chat != nil {
		msg.Chat = &Chat{
			ID:        message.Chat.ID,
			Type:      message.Chat.Type,
			Title:     message.Chat.Title,
			UserName:  message.Chat.UserName,
			FirstName: message.Chat.FirstName,
			LastName:  message.Chat.LastName,
		}
	}

	return msg
}

func getMessageType(message *tgbotapi.Message) MessageType {
	switch {
	case message.Text != "":
		return TextMessage
	case message.Photo != nil:
		return PhotoMessage
	case message.Video != nil:
		return VideoMessage
	case message.Document != nil:
		return DocumentMessage
	case message.Audio != nil:
		return AudioMessage
	case message.Voice != nil:
		return VoiceMessage
	case message.Sticker != nil:
		return StickerMessage
	default:
		return TextMessage
	}
}
