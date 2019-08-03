package orders

import (
	"encoding/json"
	"net/http"

	"github.com/brucebales/discounts-api/src/internal/discounts"
	"github.com/brucebales/discounts-api/src/internal/dto"
)

//GetOrder takes an HTTP request and attempts to populate the Order struct with it.
func GetOrder(r *http.Request, p []dto.Product) (Result, error) {
	//Scan request values into blank instance of Order
	o := Order{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&o)
	if err != nil {
		return Result{}, err
	}

	processed := o.ProcessOrder(p)

	return processed, nil
}

//ProcessOrder uses values from an order and finds matching discounts
func (o *Order) ProcessOrder(p []dto.Product) Result {
	//discountMap contains a list of discounts and their Euro values
	var discountMap = make(map[string]float64)
	//total is a copy of the Order's total, so that the original order total is preserved
	var total = o.Total

	//Slice containing tools to send to tool discount function
	var tools []dto.Item
	var switches []dto.Item

	//For loop sorts items into categories
	for _, i := range o.Items {
		for _, p := range p {
			if i.ProductID == p.ID {
				//If category is 1, append to tools
				if p.Category == 1 {
					tools = append(tools, i)
				}
				//If category is 2, append to switches
				if p.Category == 2 {
					switches = append(switches, i)
				}
			}
		}
	}

	//Feed tools slice into toolDiscount, if result is over 0 subtract discount
	if t, _ := discounts.ToolDiscount(tools); t > 0 {
		toolDiscount, reason := discounts.ToolDiscount(tools)
		total = total - toolDiscount
		discountMap[reason] = toolDiscount
	}

	//Feed switch slice into switchDiscount, if result is over 0 subtract discount
	if s, _ := discounts.SwitchDiscount(switches); s > 0 {
		switchDiscount, reason := discounts.SwitchDiscount(switches)
		total = total - switchDiscount
		discountMap[reason] = switchDiscount
	}

	// 10% off any order over 1.000 EUR
	if o.Total > 1000 {
		total = total - (o.Total * 0.10)
		discountMap["TenPercentOverOneThousand"] = o.Total - (o.Total * 0.10)
	}

	//Create JSON string of discount list
	dbytes, err := json.Marshal(discountMap)
	if err != nil {
		return Result{}
	}
	dstring := string(dbytes)

	//Create JSON string of original order
	obytes, err := json.Marshal(o)
	if err != nil {
		return Result{}
	}
	ostring := string(obytes)

	//Return final result
	return Result{
		Order:     ostring,
		Discounts: dstring,
		Total:     total,
	}
}
