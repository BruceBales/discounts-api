package dto

import (
	"database/sql"
	"encoding/json"
)

/*
[
    {
      "id": "A101",
      "description": "Screwdriver",
      "category": "1",
      "price": "9.75"
    },
    {
      "id": "A102",
      "description": "Electric screwdriver",
      "category": "1",
      "price": "49.50"
    },
    {
      "id": "B101",
      "description": "Basic on-off switch",
      "category": "2",
      "price": "4.99"
    },
    {
      "id": "B102",
      "description": "Press button",
      "category": "2",
      "price": "4.99"
    },
    {
      "id": "B103",
      "description": "Switch with motion detector",
      "category": "2",
      "price": "12.95"
    }
  ]
*/

type Product struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Category    int     `json:"category,string"`
	Price       float64 `json:"price,string"`
}

//GetProducts populates the product list from a MySQL DB entry
func GetProducts(m *sql.DB) ([]Product, error) {
	var productstr string
	products := []Product{}
	rows, err := m.Query("SELECT product_list FROM webstore.product_sets WHERE id = 1")
	if err != nil {
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&productstr)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(productstr), &products)
		if err != nil {
			return nil, err
		}
	}

	return products, nil
}
