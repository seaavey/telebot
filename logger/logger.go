package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String returns the string representation of a LogLevel
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger represents a logger instance
type Logger struct {
	level   LogLevel
	logger  *log.Logger
	service string
}

// New creates a new logger instance
func New(service string) *Logger {
	return &Logger{
		level:   INFO,
		logger:  log.New(os.Stdout, "", 0),
		service: service,
	}
}

// NewWithLevel creates a new logger instance with a specific log level
func NewWithLevel(service string, level LogLevel) *Logger {
	return &Logger{
		level:   level,
		logger:  log.New(os.Stdout, "", 0),
		service: service,
	}
}

// SetLevel sets the minimum log level for the logger
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a message at DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.log(DEBUG, fmt.Sprint(v...))
	}
}

// Debugf logs a formatted message at DEBUG level
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.log(DEBUG, fmt.Sprintf(format, v...))
	}
}

// Info logs a message at INFO level
func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.log(INFO, fmt.Sprint(v...))
	}
}

// Infof logs a formatted message at INFO level
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		l.log(INFO, fmt.Sprintf(format, v...))
	}
}

// Warn logs a message at WARN level
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.log(WARN, fmt.Sprint(v...))
	}
}

// Warnf logs a formatted message at WARN level
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		l.log(WARN, fmt.Sprintf(format, v...))
	}
}

// Error logs a message at ERROR level
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.log(ERROR, fmt.Sprint(v...))
	}
}

// Errorf logs a formatted message at ERROR level
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.log(ERROR, fmt.Sprintf(format, v...))
	}
}

// Fatal logs a message at FATAL level and exits the program
func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf logs a formatted message at FATAL level and exits the program
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(FATAL, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// log writes a log message with the specified level
func (l *Logger) log(level LogLevel, message string) {
	// Get caller information
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	} else {
		// Trim the file path to be relative to the project
		split := strings.Split(file, "telebot/")
		if len(split) > 1 {
			file = split[1]
		}
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] [%s] [%s] %s:%d %s", 
		timestamp, level.String(), l.service, file, line, message)
	
	l.logger.Println(logMessage)
}