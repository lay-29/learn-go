package model

type Order struct {
	ID     int64   `json:"id"`
	UserID int64   `json:"user_id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}
