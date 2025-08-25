package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"time"
)

func FetchJSON(url string, target interface{}) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func GetContentType(url string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	
	resp, err := client.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("bad status: " + resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", errors.New("content-type header not found")
	}

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", err
	}

	return mediaType, nil
}