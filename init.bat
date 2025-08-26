@echo off

echo Initializing telebot project...

where go >nul 2>&1
if %errorlevel% neq 0 (
    echo Go is not installed. Please install Go first: https://golang.org/doc/install
    exit /b 1
)

where docker >nul 2>&1
if %errorlevel% neq 0 (
    echo Docker is not installed. Please install Docker: https://docs.docker.com/get-docker/
    exit /b 1
)

if not exist .env (
    echo Creating .env file from .env.example...
    copy .env.example .env >nul
    echo Please update the .env file with your actual Telegram bot token and other configurations.
) else (
    echo .env file already exists.
)

echo Installing Go dependencies...
go mod tidy

echo Building the project...
go build -o telebot.exe .

if %errorlevel% equ 0 (
    echo Build successful!
    del telebot.exe
) else (
    echo Build failed. Please check for errors.
    exit /b 1
)

echo Building Docker image...
docker build -t telebot .

if %errorlevel% equ 0 (
    echo Docker image built successfully!
) else (
    echo Docker build failed. Please check for errors.
    exit /b 1
)

echo Initialization complete!
echo To run the bot, use: update.sh