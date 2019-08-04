package orders

import (
	"math"
	"net/http/httptest"
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

func TestProcessOrder_general(t *testing.T) {
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
		},
		Total: 226.2,
	}

	expected := Result{
		Order:     &Order{},
		Discounts: map[string]float64{},
		Total:     203.35,
	}

	processed, err := order.ProcessOrder(products)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	res, err := processed.String()
	if err != nil {
		t.Error("Could not create result string: ", err)
	}

	exp, err := expected.String()
	if err != nil {
		t.Error("Could not create result string: ", err)
	}

	if exp != res {
		t.Errorf("Expected: %v\n Actual: %v\n", expected, processed)
	}
}

// func TestProcessOrder_overOneThousand(t *testing.T) {
// 	products := []dto.Product{
// 		{
// 			ID:          "B103",
// 			Description: "Switch with motion detector",
// 			Category:    2,
// 			Price:       12.95,
// 		},
// 	}

// 	order := Order{
// 		ID:         "1",
// 		CustomerID: "1",
// 		Items: []dto.Item{
// 			{
// 				ProductID: "B103",
// 				Quantity:  1,
// 				UnitPrice: 2000,
// 				Total:     2000,
// 			},
// 		},
// 		Total: 2000,
// 	}

// 	expected := Result{
// 		Order:     `{"id":"1","customer-id":"1","items":[{"product-id":"B103","quantity":"1","unit-price":"2000","total":"2000"}],"total":"2000"}`,
// 		Discounts: `{"TenPercentOverOneThousand":1800}`,
// 		Total:     1800,
// 	}

// 	processed, err := order.ProcessOrder(products)
// 	if err != nil {
// 		t.Error("Unexpected error: ", err)
// 	}
// 	if expected != processed {
// 		t.Errorf("Expected: %v\n Actual: %v\n", expected, processed)
// 	}
// }

func TestProcessOrder_error(t *testing.T) {
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
		},
		{
			ID:          "A102",
			Description: "Electric screwdriver",
			//Mess up json Marshaling with infinite number
			Price: math.Inf(1),
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
				//Mess up json Marshaling with infinite number
				Total: math.Inf(1),
			},
			{
				ProductID: "A102",
				Quantity:  3,
				UnitPrice: 49.50,
				Total:     148.5,
			},
		},
		//Mess up json Marshaling with infinite number
		Total: math.Inf(1),
	}

	_, err := order.ProcessOrder(products)
	if err == nil {
		t.Error("Expected error, but recieved none")
	}
}

func TestProcessOrder_error_originalOrder(t *testing.T) {
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
		},
		{
			ID:          "A102",
			Description: "Electric screwdriver",
			Price:       1337,
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
				//Mess up json Marshaling with infinite number
				Total: math.Inf(1),
			},
			{
				ProductID: "A102",
				Quantity:  3,
				UnitPrice: 49.50,
				Total:     148.5,
			},
		},
		Total: 1337,
	}

	_, err := order.ProcessOrder(products)
	if err == nil {
		t.Error("Expected error, but recieved none")
	}
}
