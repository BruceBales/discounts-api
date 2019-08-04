package orders

import (
	"math"
	"testing"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

func TestString(t *testing.T) {
	//Mock result to stringify
	result := Result{
		Order: &Order{
			ID:         "1",
			CustomerID: "1",
			Items: []dto.Item{
				{
					ProductID: "B103",
					Quantity:  6,
					UnitPrice: 12.95,
					Total:     77.7,
				},
				{
					ProductID: "A102",
					Quantity:  3,
					UnitPrice: 49.50,
					Total:     148.5,
				},
				{
					ProductID: "C123",
					Quantity:  3,
					UnitPrice: 1000,
					Total:     3000,
				},
			},
			Total: 3226.2,
		},
		Discounts: map[string]float64{
			"FirstToolDiscount":         9.9,
			"TenPercentOverOneThousand": 2903.58,
			"switchDiscount":            12.95,
		},
		Total: 2880.73,
	}
	//Call string function to stringify result
	_, err := result.String()
	if err != nil {
		t.Error("Could not convert result to string: ", err)
	}

}

func TestString_error(t *testing.T) {
	//Mock result to stringify
	result := Result{
		Order: &Order{
			ID:         "1",
			CustomerID: "1",
			Items: []dto.Item{
				{
					ProductID: "B103",
					Quantity:  6,
					UnitPrice: 12.95,
					Total:     77.7,
				},
				{
					ProductID: "A102",
					Quantity:  3,
					UnitPrice: 49.50,
					Total:     148.5,
				},
				{
					ProductID: "C123",
					Quantity:  3,
					UnitPrice: 1000,
					Total:     3000,
				},
			},
			Total: 3226.2,
		},
		Discounts: map[string]float64{
			"FirstToolDiscount":         9.9,
			"TenPercentOverOneThousand": 2903.58,
			"switchDiscount":            12.95,
		},
		//Curveball to force json.Marshal to mess up
		Total: math.Inf(1),
	}

	_, err := result.String()
	if err == nil {
		t.Error("Expected error but received none")
	}

}
