package orders

import (
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
		Order:     `{"id":"1","customer-id":"1","items":[{"product-id":"B103","quantity":"6","unit-price":"12.95","total":"77.7"},{"product-id":"A102","quantity":"3","unit-price":"49.5","total":"148.5"}],"total":"226.2"}`,
		Discounts: `{"FirstToolDiscount":9.9,"switchDiscount":12.95}`,
		Total:     203.35,
	}

	processed := order.ProcessOrder(products)

	if expected != processed {
		t.Errorf("Expected: %v\n Actual: %v\n", expected, processed)
	}

}

func TestProcessOrder_overOneThousand(t *testing.T) {
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
			Price:       12.95,
		},
	}

	order := Order{
		ID:         "1",
		CustomerID: "1",
		Items: []dto.Item{
			{
				ProductID: "B103",
				Quantity:  1,
				UnitPrice: 2000,
				Total:     2000,
			},
		},
		Total: 2000,
	}

	expected := Result{
		Order:     `{"id":"1","customer-id":"1","items":[{"product-id":"B103","quantity":"1","unit-price":"2000","total":"2000"}],"total":"2000"}`,
		Discounts: `{"TenPercentOverOneThousand":1800}`,
		Total:     1800,
	}

	processed := order.ProcessOrder(products)

	if expected != processed {
		t.Errorf("Expected: %v\n Actual: %v\n", expected, processed)
	}

}
