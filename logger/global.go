package logger

import (
	"os"
	"strconv"
)

// Global logger instance
var defaultLogger *Logger

func init() {
	// Initialize default logger
	defaultLogger = New("telebot")
	
	// Set log level from environment variable, default to INFO
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		switch logLevel {
		case "DEBUG":
			defaultLogger.SetLevel(DEBUG)
		case "INFO":
			defaultLogger.SetLevel(INFO)
		case "WARN":
			defaultLogger.SetLevel(WARN)
		case "ERROR":
			defaultLogger.SetLevel(ERROR)
		case "FATAL":
			defaultLogger.SetLevel(FATAL)
		}
	}
	
	// Check if debug mode is enabled
	if debugMode := os.Getenv("BOT_DEBUG"); debugMode != "" {
		if debug, err := strconv.ParseBool(debugMode); err == nil && debug {
			defaultLogger.SetLevel(DEBUG)
		}
	}
}

// Default returns the default logger instance
func Default() *Logger {
	return defaultLogger
}

// Debug logs a message at DEBUG level using the default logger
func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

// Debugf logs a formatted message at DEBUG level using the default logger
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Info logs a message at INFO level using the default logger
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Infof logs a formatted message at INFO level using the default logger
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Warn logs a message at WARN level using the default logger
func Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

// Warnf logs a formatted message at WARN level using the default logger
func Warnf(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v...)
}

// Error logs a message at ERROR level using the default logger
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

// Errorf logs a formatted message at ERROR level using the default logger
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Fatal logs a message at FATAL level using the default logger and exits
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

// Fatalf logs a formatted message at FATAL level using the default logger and exits
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}