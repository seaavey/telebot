# Telebot - Multi-Platform Media Downloader Bot

A Telegram bot built with Go that allows users to download media content from popular platforms without watermarks.

## Features
- Download TikTok videos and images without watermarks
- Download Instagram photos and videos
- Download Pinterest images
- Access MediaFire files with direct download links
- Simple command interface with `/start` command
- Automatic URL detection in messages

## Supported Platforms
- TikTok
- Instagram
- Pinterest
- MediaFire

## How to Use
1. Start the bot with `/start`
2. Send a link from any supported platform
3. Receive the media content directly in Telegram

## Technical Details
- Built with Go using the `go-telegram-bot-api` library
- Uses external APIs for media processing
- Modular architecture with handlers, services, and utilities

## Setup
1. Clone the repository
2. Create a `.env` file with your Telegram bot token:
   ```
   TELEGRAM_BOT_TOKEN=your_bot_token_here
   ```
3. Run `go mod tidy` to install dependencies
4. Start the bot with `go run main.go`

## Dependencies
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [godotenv](https://github.com/joho/godotenv)