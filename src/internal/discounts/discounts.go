package discounts

import (
	"sort"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

func SwitchDiscount(switches []dto.Item) (float64, string) {

	var discount float64

	//I am making a huge assumption that this discount is per unique switch item
	for _, i := range switches {
		if i.Quantity > 5 {
			discount += i.UnitPrice
		}
	}

	if discount > 2 {
		return discount, "switchDiscount"
	}

	return 0, ""
}

func ToolDiscount(tools []dto.Item) (float64, string) {
	var total int
	//Count tools
	for _, i := range tools {
		total = total + i.Quantity
	}
	if total > 2 {
		//Create array of all prices
		var prices []float64
		for _, tool := range tools {
			prices = append(prices, tool.UnitPrice)
		}
		//Sort array of prices to find lowest
		sort.Float64s(prices)

		//Grab price from first tool in sorted array, return discount
		price := prices[0]
		discount := price * 0.20
		return discount, "FirstToolDiscount"
	}
	//No discount if less than two tools
	return 0, ""
}
