package types

type PinterestResponse struct {
	Creator string `json:"creator"`
	Status  int    `json:"status"`
	Data    struct {
		Thumbnail string `json:"thumbnail"`
		Url       string `json:"url"`
		Title     string `json:"title"`
	} `json:"data"`
}