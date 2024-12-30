package models

type Order struct {
	ID         int     `json:"id"`
	ClientID   int     `json:"client_id"`
	ProductIDs []int   `json:"product_ids"`
	Total      float64 `json:"total"`
}
