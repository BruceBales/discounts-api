package orders

import (
	"encoding/json"
	"net/http"
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
	ID         string  `json:"id"`
	CustomerID string  `json:"customer-id"`
	Items      []Item  `json:"items"`
	Total      float64 `json:"total,string"`
}

//Item is a blueprint for each orderable item in the store
type Item struct {
	ProductID string  `json:"product-id"`
	Quantity  int     `json:"quantity,string"`
	UnitPrice float64 `json:"unit-price,string"`
	Total     float64 `json:"total,string"`
}

type Result struct {
	Order     string
	Discounts string
	Total     float64
}

//GetOrder takes an HTTP request and attempts to populate the Order struct with it.
func GetOrder(r *http.Request, p []Product) (Result, error) {
	o := Order{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&o)
	if err != nil {
		return Result{}, err
	}

	processed := o.ProcessOrder(p)

	return processed, nil
}

func (o *Order) ProcessOrder(p []Product) Result {
	/*This function is currently a huge mess. Will need to refactor
	once base functionality is working.*/

	var discounts = make(map[string]float64)

	var total = o.Total

	for _, i := range o.Items {
		for _, p := range p {
			if i.ProductID == p.ID {
				if p.Category == 2 && i.Quantity > 5 {
					//Give one extra for free
					i.Quantity++
					discounts["Free6thSwitch"] = i.UnitPrice
					total = total - i.UnitPrice
				}
				if p.Category == 1 && i.Quantity > 2 {
					//Make first one free
					total = total - i.UnitPrice
					discounts["FirstToolFree"] = i.UnitPrice
				}
			}
		}
	}

	if o.Total > 1000 {
		total = o.Total * 0.10
	}

	dbytes, err := json.Marshal(discounts)
	if err != nil {
		return Result{}
	}
	dstring := string(dbytes)

	obytes, err := json.Marshal(o)
	if err != nil {
		return Result{}
	}
	ostring := string(obytes)

	return Result{
		Order:     ostring,
		Discounts: dstring,
		Total:     total,
	}
}
