package orders

import (
	"encoding/json"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

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
	Order     *Order
	Discounts map[string]float64
	Total     float64
}

//String returns a JSON string of the Result struct
func (r *Result) String() (string, error) {
	//Create byte slice of Marshal'ed result
	resBytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	//Type-cast byte slice to string
	str := string(resBytes)
	return str, nil
}
