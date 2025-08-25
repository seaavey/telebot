package types

type TikTokResponse struct {
	Creator string `json:"creator"`
	Status  int    `json:"status"`
	Data    struct {
		ID        float64 `json:"id"`
		Title     string  `json:"title"`
		URL       string  `json:"url"`
		CreatedAt string  `json:"created_at"`

		Images []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`

		Video struct {
			NoWatermark string `json:"noWatermark"`
		} `json:"videos"`
	} `json:"data"`
}