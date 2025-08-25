# Telebot - Multi-Platform Media Downloader Bot

A Telegram bot built with Go that allows users to download media content from popular platforms without watermarks.

## Features
- Download TikTok videos and images without watermarks
- Download Instagram photos and videos
- Download Pinterest images
- Download Facebook videos and images
- Simple command interface with `/start` command
- Automatic URL detection in messages
- Improved error handling and timeouts for HTTP requests
- Better validation and error reporting
- Structured logging with multiple log levels

## Supported Platforms
- TikTok
- Instagram
- Pinterest
- Facebook

## How to Use
1. Start the bot with `/start`
2. Send a link from any supported platform
3. Receive the media content directly in Telegram

## Technical Details
- Built with Go using the `go-telegram-bot-api` library
- Uses external APIs for media processing
- Modular architecture with handlers, services, and utilities
- Enhanced HTTP client with 30-second timeouts
- Improved error handling with contextual error messages
- Custom logger with multiple log levels (DEBUG, INFO, WARN, ERROR, FATAL)
- Environment-based configuration for log levels and debug mode

## Setup
1. Clone the repository
2. Create a `.env` file with your Telegram bot token:
   ```
   TELEGRAM_BOT_TOKEN=your_bot_token_here
   BOT_DEBUG=false
   LOG_LEVEL=INFO
   ```
3. Run `go mod tidy` to install dependencies
4. Start the bot with `go run main.go`

## Configuration
- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token (required)
- `BOT_DEBUG`: Enable debug mode (true/false, default: false)
- `LOG_LEVEL`: Set log level (DEBUG, INFO, WARN, ERROR, FATAL, default: INFO)
- `BOT_OWNERS`: Comma-separated list of Telegram user IDs that are bot owners

## Recent Improvements
- **Facebook Downloader**: Added support for downloading Facebook videos and images
- **Enhanced Logging**: Implemented a custom logger with multiple log levels and caller information
- **Refactored Platform Detection**: Improved URL processing with switch statements for better readability
- **Enhanced HTTP Client**: Added 30-second timeouts to all HTTP requests to prevent hanging
- **Better Error Handling**: Improved error messages with context for easier debugging
- **URL Validation**: Fixed URL extraction to properly report when no valid URL is found

## Error Handling
The bot implements comprehensive error handling for all supported platforms. When an error occurs, users will receive a descriptive message explaining the issue rather than a generic failure notification.

## Dependencies
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [godotenv](https://github.com/joho/godotenv)