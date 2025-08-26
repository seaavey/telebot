#!/bin/bash

CONTAINER_NAME="my-telebot"

if [ $# -eq 1 ]; then
    CONTAINER_NAME="$1"
fi

if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first: https://golang.org/doc/install"
    exit 1
fi

if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Please install Docker: https://docs.docker.com/get-docker/"
    exit 1
fi

if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "Please update the .env file with your actual Telegram bot token and other configurations."
else
    echo ".env file already exists."
fi

echo "Installing Go dependencies..."
go mod tidy

echo "Building the project..."
go build -o telebot .

if [ $? -eq 0 ]; then
    echo "Build successful!"
    rm telebot
else
    echo "Build failed. Please check for errors."
    exit 1
fi

echo "Building Docker image..."
docker build -t telebot .

if [ $? -eq 0 ]; then
    echo "Docker image built successfully!"
else
    echo "Docker build failed. Please check for errors."
    exit 1
fi

echo "Initialization complete!"
echo "To run the bot, use: ./update.sh"