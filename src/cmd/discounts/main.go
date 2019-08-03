package main

import (
	"fmt"
	"net/http"

	"github.com/brucebales/discounts-api/src/internal/dto"
	"github.com/brucebales/discounts-api/src/internal/orders"
)

func main() {
	//Establish redis connection
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: ":6379",
	// })

	//Hard-coding for testing but will have to replace with redis query
	products := []dto.Product{
		{
			ID:          "B103",
			Description: "Switch with motion detector",
			Category:    2,
			Price:       12.95,
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Invalid path, please read API documentation")
	})

	http.HandleFunc("/api/order", func(w http.ResponseWriter, r *http.Request) {
		result, err := orders.GetOrder(r, products)
		if err != nil {
			fmt.Println("Could not get order: ", err)
		}
		fmt.Fprintf(w, fmt.Sprintf("Order: %v\n", result.Order))
		fmt.Fprintf(w, fmt.Sprintf("Discounts: %v\n", result.Discounts))
		fmt.Fprintf(w, fmt.Sprintf("Total: %v\n", result.Total))
	})

	fmt.Println("Discounts API running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Could not serve HTTP")
	}
}
