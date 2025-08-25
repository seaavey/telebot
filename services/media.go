package services

import (
	"fmt"
	"strings"
	"telebot/types"
	"telebot/utils"
)

const (
	TikTokAPIURL    = "https://api.seaavey.my.id/api/downloader/tiktok?url="
	InstagramAPIURL = "https://api.seaavey.my.id/api/downloader/instagram?url="
	MediaFireURL    = "https://api.seaavey.my.id/api/downloader/mediafire?url="
	PinterestAPIURL = "https://api.seaavey.my.id/api/downloader/pinterest?url="
)

type MediaService struct{}

func NewMediaService() *MediaService {
	return &MediaService{}
}

func (s *MediaService) ProcessMediaFireURL(url string) (*types.MediafireResponse, error) {
	var data types.MediafireResponse
	err := utils.FetchJSON(MediaFireURL+url, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mediafire data: %w", err)
	}

	if data.Status != 200 {
		return nil, fmt.Errorf("API returned status %d", data.Status)
	}

	return &data, nil
}
func (s *MediaService) ProcessTikTokURL(url string) (*types.TikTokResponse, error) {
	var data types.TikTokResponse
	err := utils.FetchJSON(TikTokAPIURL+url, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch TikTok data: %w", err)
	}

	if data.Status != 200 {
		return nil, fmt.Errorf("API returned status %d", data.Status)
	}

	return &data, nil
}

func (s *MediaService) ProcessInstagramURL(url string) (*types.InstagramResponse, error) {
	var data types.InstagramResponse
	err := utils.FetchJSON(InstagramAPIURL+url, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Instagram data: %w", err)
	}

	if data.Status != 200 {
		return nil, fmt.Errorf("API returned status %d", data.Status)
	}

	return &data, nil
}

func (s *MediaService) ProcessPinterestURL(url string) (*types.PinterestResponse, error) {
	var data types.PinterestResponse
	err := utils.FetchJSON(PinterestAPIURL+url, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Pinterest data: %w", err)
	}

	if data.Status != 200 {
		return nil, fmt.Errorf("API returned status %d", data.Status)
	}

	return &data, nil
}

func (s *MediaService) IsTikTokURL(url string) bool {
	return strings.Contains(url, "tiktok.com")
}

func (s *MediaService) IsInstagramURL(url string) bool {
	return strings.Contains(url, "instagram.com") || strings.Contains(url, "instagr.am")
}

func (s *MediaService) IsMediaFireURL(url string) bool {
	return strings.Contains(url, "mediafire.com")
}

func (s *MediaService) IsPinterestURL(url string) bool {
	return strings.Contains(url, "pinterest.com") || strings.Contains(url, "pin.it")
}