#!/bin/bash

CONTAINER_NAME="my-telebot"

if [ $# -eq 1 ]; then
    CONTAINER_NAME="$1"
fi

if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Installing Docker..."
    
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get &> /dev/null; then
            sudo apt-get update
            sudo apt-get install -y docker.io
            sudo systemctl start docker
            sudo systemctl enable docker
        elif command -v yum &> /dev/null; then
            sudo yum install -y docker
            sudo systemctl start docker
            sudo systemctl enable docker
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y docker
            sudo systemctl start docker
            sudo systemctl enable docker
        fi
        
        sudo usermod -aG docker $USER
        echo "Docker installed. Please log out and back in to use Docker without sudo."
    else
        echo "Unsupported OS. Please install Docker manually."
        exit 1
    fi
fi

cd "$(dirname "$0")"

git pull

docker build -t telebot .

docker stop "$CONTAINER_NAME" 2>/dev/null || echo "Container $CONTAINER_NAME not running or doesn't exist"
docker rm "$CONTAINER_NAME" 2>/dev/null || echo "Container $CONTAINER_NAME not found"

docker run -dit --name "$CONTAINER_NAME" --env-file .env telebot

echo "Container $CONTAINER_NAME updated and running!"
