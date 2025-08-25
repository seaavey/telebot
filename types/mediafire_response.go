package types

type MediafireResponse struct {
	Creator string `json:"creator"`
	Status  int    `json:"status"`
	Data    struct {
		Filename string `json:"filename"`
		Dl       string `json:"dl"`
	} `json:"data"`
}