package models

type Product struct {
	ID    int64 `json:"id"`
	Count int64 `json:"count"`
}

type Cart struct {
	UserID   int64     `json:"user_id"`
	Products []*Product `json:"products"`
}
