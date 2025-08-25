package types

type InstagramResponse struct {
	Creator string   `json:"creator"`
	Status  int      `json:"status"`
	Data    []string `json:"data"`
}