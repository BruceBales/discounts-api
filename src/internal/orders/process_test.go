package orders

import (
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

func TestGetOrder_general(t *testing.T) {
	//Mock product list, shorter than the real one
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
			Price:       12.95,
		},
		{
			ID:          "A102",
			Description: "Electric screwdriver",
			Category:    1,
			Price:       49.50,
		},
	}

	//Mock POST body
	body := strings.NewReader(`{
		"id": "1",
		"customer-id": "1",
		"items": [{
			"product-id": "B103",
			"quantity": "6",
			"unit-price": "4.99",
			"total": "49.90"
		}],
		"total": "49.90"
	}`)

	//Mock request using post body
	r := httptest.NewRequest("POST", "/api/discounts", body)

	//Call to GetOrder
	_, err := GetOrder(r, products)
	if err != nil {
		t.Errorf("Error parsing request: %s", err)
	}
}

func TestGetOrder_error(t *testing.T) {
	//Mock product list
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
			Price:       12.95,
		},
		{
			ID:          "A102",
			Description: "Electric screwdriver",
			Category:    1,
			Price:       49.50,
		},
	}
	//Mock POST body with missing comma in JSON to trigger error
	body := strings.NewReader(`{
		"id": "1",
		"customer-id": "1",
		"items": [{
			"product-id": "B103",
			"quantity": "6",
			"unit-price": "4.99"
			"total": "49.90"
		}],
		"total": "49.90"
	}`)

	//Mock HTTP request
	r := httptest.NewRequest("POST", "/api/discounts", body)

	//Call to GetOrder that expects an error due to invalid JSON
	_, err := GetOrder(r, products)
	if err.Error() != `invalid character '"' after object key:value pair` {
		t.Error("Failed to catch invalid JSON")
	}
}

func TestProcessOrder(t *testing.T) {
	//Mock product list
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
			Price:       12.95,
		},
		{
			ID:          "A102",
			Description: "Electric screwdriver",
			Category:    1,
			Price:       49.50,
		},
	}

	//Mock order, in struct form
	order := Order{
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
	}

	//Expected result after calculating discounts
	expected := Result{
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

	//Get actual results through ProcessOrder call
	processed := order.ProcessOrder(products)

	//Use DeepEqual from reflect package to evaluate results,
	//because we have a map in the Result struct
	match := reflect.DeepEqual(expected, processed)
	if !match {
		//Getting strings of results, only for use in debugging if test fails
		exp, err := expected.String()
		if err != nil {
			t.Error("Error converting expected to string: ", err)
		}
		proc, err := processed.String()
		if err != nil {
			t.Error("Error converting actual to string: ", err)
		}

		t.Errorf("Expected: %v\n Actual: %v\n", exp, proc)
	}
}
