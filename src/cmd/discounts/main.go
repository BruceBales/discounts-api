package main

import (
	"fmt"
	"net/http"

	"github.com/brucebales/discounts-api/src/internal/dao"
	"github.com/brucebales/discounts-api/src/internal/dto"
	"github.com/brucebales/discounts-api/src/internal/orders"
)

func main() {
	//Establish Redis connection
	// redis, err := dao.NewRedis()
	// if err != nil {
	// 	fmt.Println("Could not connect to Redis: ", err)
	// }
	//Establish MySQL connection
	mysql, err := dao.NewMysql()
	if err != nil {
		fmt.Println("Coult not connect to MySQL: ", err)
	}

	//Handle HTTP Requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Invalid path")
	})

	http.HandleFunc("/api/order", func(w http.ResponseWriter, r *http.Request) {
		products, err := dto.GetProducts(mysql)
		if err != nil {
			fmt.Println("Could not load products: ", err)
		}
		fmt.Println(products)
		result, err := orders.GetOrder(r, products)
		if err != nil {
			fmt.Println("Could not get order: ", err)
		}
		fmt.Fprintf(w, fmt.Sprintf("Order: %v\n", result.Order))
		fmt.Fprintf(w, fmt.Sprintf("Discounts: %v\n", result.Discounts))
		fmt.Fprintf(w, fmt.Sprintf("Total: %v\n", result.Total))
	})

	fmt.Println("Discounts API running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Could not serve HTTP")
	}
}
