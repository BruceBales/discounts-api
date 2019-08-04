package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brucebales/discounts-api/src/internal/dao"
	"github.com/brucebales/discounts-api/src/internal/dto"
	"github.com/brucebales/discounts-api/src/internal/orders"
)

func main() {
	//Establish Redis connection
	redis, err := dao.NewRedis()
	if err != nil {
		fmt.Println("Could not connect to Redis: ", err)
		return
	}
	//Establish MySQL connection
	mysql, err := dao.NewMysql()
	if err != nil {
		fmt.Println("Coult not connect to MySQL: ", err)
		return
	}
	defer mysql.Close()

	//- Handle HTTP Requests -
	//Return 404 not found for root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Invalid path")
	})

	//Handle Orders
	http.HandleFunc("/api/order", func(w http.ResponseWriter, r *http.Request) {
		//Fetch products on each request to allow updating product list live
		products, err := dto.GetProducts(mysql)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error loading products")
			fmt.Println("Could not load products: ", err)
			return
		}

		//Fetch results by inputting current product list and HTTP request into GetOrder
		result, err := orders.GetOrder(r, products)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error processing order")
			fmt.Println("Could not get order: ", err)
			return
		}

		//Create response JSON string
		response, err := result.String()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error creating response")
			fmt.Println("Could not return response: ", err)
			return
		}
		//Log discounts to Redis
		redis.SAdd("discount_log", fmt.Sprintf(`[{"time":"%s","result":%s}]`, time.Now(), response))

		//Print response
		fmt.Fprint(w, response)
	})

	fmt.Println("Discounts API running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Could not serve HTTP")
	}
}
