package dto

type Item struct {
	ProductID string  `json:"product-id"`
	Quantity  int     `json:"quantity,string"`
	UnitPrice float64 `json:"unit-price,string"`
	Total     float64 `json:"total,string"`
}
