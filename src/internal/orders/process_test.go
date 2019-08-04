package orders

import (
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/brucebales/discounts-api/src/internal/dto"
)

func TestGetOrder_general(t *testing.T) {
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

	r := httptest.NewRequest("POST", "/api/discounts", body)

	_, err := GetOrder(r, products)
	if err != nil {
		t.Errorf("Error parsing request: %s", err)
	}
}

func TestGetOrder_error(t *testing.T) {
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

	r := httptest.NewRequest("POST", "/api/discounts", body)

	_, err := GetOrder(r, products)
	if err.Error() != `invalid character '"' after object key:value pair` {
		t.Error("Failed to catch invalid JSON")
	}
}

func TestProcessOrder(t *testing.T) {
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

	processed := order.ProcessOrder(products)

	match := reflect.DeepEqual(expected, processed)

	exp, err := expected.String()
	if err != nil {
		t.Error("Error converting expected to string: ", err)
	}
	proc, err := processed.String()
	if err != nil {
		t.Error("Error converting actual to string: ", err)
	}

	if !match {
		t.Errorf("Expected: %v\n Actual: %v\n", exp, proc)
	}
}
