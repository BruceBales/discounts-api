package orders

import "github.com/brucebales/discounts-api/src/internal/dto"

/*
Example order:
{
    "id": "1",
    "customer-id": "1",
    "items": [
      {
        "product-id": "B102",
        "quantity": "10",
        "unit-price": "4.99",
        "total": "49.90"
      }
    ],
    "total": "49.90"
  }
*/

//Order defines values accepted in an Order request
type Order struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customer-id"`
	Items      []dto.Item `json:"items"`
	Total      float64    `json:"total,string"`
}

//Result is returned to main as the final copy of all processed data
type Result struct {
	Order     string
	Discounts string
	Total     float64
}
