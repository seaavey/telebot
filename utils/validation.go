package utils

import (
	"net/url"
	"regexp"
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
		return "", nil
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