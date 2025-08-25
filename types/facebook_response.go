package types

type FacebookResponse struct {
	Status    bool   `json:"status"`
	Timestamp string `json:"timestamp"`
	Data      struct {
		Thumbnail string `json:"thumbnail"`
		Title     string `json:"title"`
		Data      []struct {
			URL        string `json:"url"`
			Resolution string `json:"resolution"`
			Format     string `json:"format"`
		} `json:"data"`
	} `json:"data"`
}