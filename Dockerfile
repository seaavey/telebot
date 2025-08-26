# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o telebot .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -s /bin/sh telebot

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/telebot .

# Change ownership of the binary to the non-root user
RUN chown -R telebot:telebot /app

# Switch to the non-root user
USER telebot

# Expose port (if your app would need it in the future)
EXPOSE 8080

# Run the binary
CMD ["./telebot"]