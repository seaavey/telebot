package handlers

import (
	"log"
	"strings"
	"telebot/services"
	"telebot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	mediaService *services.MediaService
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		mediaService: services.NewMediaService(),
	}
}

func (h *MessageHandler) HandleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		h.handleMessage(update.Message, bot)
	}
}

func (handler *MessageHandler) handleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := utils.SerializeMessage(message, bot)
	
	if msg == nil {
		return
	}

	if msg.Text != "" && msg.Text[0] == '/' {
		handler.handleCommand(msg)
		return
	}

	if msg.Text != "" && utils.IsURL(msg.Text) {
		handler.processURL(msg)
	}
}

func (handler *MessageHandler) handleCommand(msg *utils.Message) {
	if msg.Text == "/start" {
		welcomeText := "👋 Welcome to the Media Downloader Bot!\n\n" +
			"I can help you download media from various platforms without watermarks.\n\n" +
			"Just send me a link from any supported platform and I'll process it for you!\n\n" +
			"Supported platforms:\n" +
			"• TikTok (videos and images)\n" +
			"• Instagram (photos and videos)\n" +
			"• Pinterest (images)\n" +
			"• Facebook (videos and images)\n\n" +
			"Note: For some platforms, I may provide a download link rather than direct media."
		
		msg.SendText(welcomeText)
		return
	}
	
	msg.SendText("❌ Sorry, I don't recognize that command.")
}

func (handler *MessageHandler) processURL(msg *utils.Message) {
	url, err := utils.GetURL(msg.Text)
	if err != nil {
		log.Printf("Error extracting URL: %v", err)
		msg.SendText("❌ Sorry, I couldn't find a valid URL in your message. Please make sure you're sending a direct link from TikTok, Instagram, Pinterest, or Facebook.")
		return
	}

	switch {
	case handler.mediaService.IsTikTokURL(url):
		handler.processTikTokURL(msg, url)
	case handler.mediaService.IsInstagramURL(url):
		handler.processInstagramURL(msg, url)
	case handler.mediaService.IsPinterestURL(url):
		handler.processPinterestURL(msg, url)
	case handler.mediaService.IsFacebookURL(url):
		handler.processFacebookURL(msg, url)
	default:
		msg.SendText("❌ Sorry, I don't support that platform yet. I can process links from TikTok, Instagram, Pinterest, and Facebook.")
	}
}

func (handler *MessageHandler) processPinterestURL(msg *utils.Message, url string) {
	data, err := handler.mediaService.ProcessPinterestURL(url)
	if err != nil {
		log.Printf("Error processing Pinterest URL: %v", err)
		msg.SendText("❌ Sorry, I couldn't process that Pinterest link.")
		return
	}

	caption := data.Data.Title
	
	contentType, err := utils.GetContentType(data.Data.Url)
	if err != nil {
		log.Printf("Error getting content type for Pinterest media: %v", err)
		err = msg.SendImage(data.Data.Url, caption)
	} else if strings.Contains(contentType, "video") {
		err = msg.SendVideo(data.Data.Url, caption)
	} else {
		err = msg.SendImage(data.Data.Url, caption)
	}
	
	if err != nil {
		log.Printf("Error sending Pinterest media: %v", err)
		msg.SendText("❌ Sorry, I couldn't send the Pinterest media.")
	}
}

func (handler *MessageHandler) processFacebookURL(msg *utils.Message, url string) {
	data, err := handler.mediaService.ProcessFacebookURL(url)
	if err != nil {
		log.Printf("Error processing Facebook URL: %v", err)
		msg.SendText("❌ Sorry, I couldn't process that Facebook link.")
		return
	}

	// Use the first media item from the response
	if len(data.Data.Data) > 0 {
		media := data.Data.Data[0]
		mediaURL := media.URL
		caption := data.Data.Title
		
		var err error
		switch {
		case media.Resolution == "Audio":
			err = msg.SendAudio(mediaURL, caption)
		case media.Format == "mp4":
			err = msg.SendVideo(mediaURL, caption)
		case media.Format == "jpg":
			err = msg.SendImage(mediaURL, caption)
		default:
			// Fallback to document for other formats
			err = msg.SendDocument(mediaURL, caption)
		}

		if err != nil {
			log.Printf("Error sending Facebook media: %v", err)
			msg.SendText("❌ Sorry, I couldn't send the Facebook media.")
		}
	} else {
		msg.SendText("❌ Sorry, I couldn't find any downloadable media in that Facebook link.")
	}
}

func (handler *MessageHandler) processTikTokURL(msg *utils.Message, url string) {
	data, err := handler.mediaService.ProcessTikTokURL(url)
	if err != nil {
		log.Printf("Error processing TikTok URL: %v", err)
		msg.SendText("❌ Sorry, I couldn't process that TikTok link.")
		return
	}

	if data.Data.Video.NoWatermark != "" {
		err = msg.SendVideo(data.Data.Video.NoWatermark, data.Data.Title)
		if err != nil {
			log.Printf("Error sending video: %v", err)
			msg.SendText("❌ Sorry, I couldn't send the video.")
		}
	} else if len(data.Data.Images) > 0 {
		for i, img := range data.Data.Images {
			var caption string
			if i == 0 {
				caption = data.Data.Title
			}
			err = msg.SendImage(img.URL, caption)
			if err != nil {
				log.Printf("Error sending image: %v", err)
				msg.SendText("❌ Sorry, I couldn't send all images.")
				break
			}
		}
	} else {
		msg.SendText("❌ Sorry, I couldn't find any video or images in that TikTok link.")
	}
}

func (handler *MessageHandler) processInstagramURL(msg *utils.Message, url string) {
	data, err := handler.mediaService.ProcessInstagramURL(url)
	if err != nil {
		log.Printf("Error processing Instagram URL: %v", err)
		msg.SendText("❌ Sorry, I couldn't process that Instagram link.")
		return
	}

	if len(data.Data) > 0 {
		for _, mediaURL := range data.Data {
			contentType, err := utils.GetContentType(mediaURL)
			if err != nil {
				log.Printf("Error getting content type for Instagram media: %v", err)
				msg.SendText("❌ Sorry, I encountered an issue processing Instagram media.")
				break
			}

			isImage := strings.Contains(contentType, "image")

			if isImage {
				err = msg.SendImage(mediaURL)
			} else {
				err = msg.SendVideo(mediaURL)
			}

			if err != nil {
				log.Printf("Error sending Instagram media: %v", err)
				msg.SendText("❌ Sorry, I couldn't send all Instagram media.")
				break
			}
		}
	} else {
		msg.SendText("❌ Sorry, I couldn't find any media in that Instagram link.")
	}
}