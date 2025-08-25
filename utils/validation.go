package utils

import (
	"errors"
	"net/url"
	"os"
	"regexp"
	"strings"
	"telebot/config"
)

var URLRegex = regexp.MustCompile(`https?://[^\s]+`)

func IsOwner(userID int64) bool {
	for _, ownerID := range config.Owners {
		if ownerID == userID {
			return true
		}
	}
	return false
}

func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func GetURL(text string) (string, error) {
	match := URLRegex.FindString(text)
	if match == "" {
		return "", errors.New("no valid URL found in text")
	}

	u, err := url.ParseRequestURI(match)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func Coalesce(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func GetEnvAsBool(name string, defaultValue bool) bool {
	valStr := os.Getenv(name)
	if valStr == "" {
		return defaultValue
	}
	valStr = strings.ToLower(valStr)

	if valStr == "true" || valStr == "1" || valStr == "on" || valStr == "enabled" {
		return true
	}
	return false
}