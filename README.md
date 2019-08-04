# Discounts API

An API for 24Metrics that consumes POST requests from a webstore and calculates discounts.

## Overview

This API takes customer orders, and applies various discounts to them. Orders are sent as a POST request with a JSON body, and results are returned as a JSON body containing the original order, a list of discounts(with monetary values), and a "total" with discounts applied to it.

The same JSON body as given in the HTTP response is also logged in a Redis database.

## Setup

To run this project, you will need the following tools installed:

- Go
- Docker
- Docker-compose
- Git

Please ensure that the project is cloned into your GOPATH under the correct namespace(github.com/brucebales/discounts-api).

### Building

To get started, first install the project's dependencies:

```
git submodule update --init
```

Next, we will build the binary into the gitignored `bin` folder:

```
go build -o ./bin/discounts ./src/cmd/discounts/main.go
```

Note: it is important to build the binary into `bin/discounts` so that Docker can access it.


### Running Locally

To run the API locally, we will use docker-compose:

```
docker-compose up
```

After the docker-compose command has been run, the API is available on port 8080. Let's send it a request:

```
curl --request POST \
  --url http://localhost:8080/api/order \
  --header 'content-type: application/json' \
  --data '{
	"id": "2",
	"customer-id": "1",
	"items": [{
		"product-id": "B103",
		"quantity": "6",
		"unit-price": "4.99",
		"total": "49.90"
	}],
	"total": "49.90"
}'
```

Example response:

```
{
  "Order": {
    "id": "2",
    "customer-id": "2",
    "items": [
      {
        "product-id": "B102",
        "quantity": "6",
        "unit-price": "4.99",
        "total": "24.95"
      }
    ],
    "total": "24.95"
  },
  "Discounts": {
    "switchDiscount": 4.99
  },
  "Total": 19.96
}
```

After this request, we can also check for logs that are stored in Redis:

```
docker-compose exec redis redis-cli
```

```
smembers discount_log
```


### Adding More Products

The product list is stored in MySQL as a JSON array string called `product_list` in `webstore.products`. To add more products, add another item to this JSON array.


## Room for improvement

A few things that I would like to complete if I end up having time:

- More extensive unit testing in all packages
- An authentication system using API keys
- Further abstraction of discounts, allowing new discounts to be defined in a database.
  - Ideally, there would be a set of "discount type" functions that are simply plugging in variables that are grabbed from a MySQL table.
- A separate endpoint for adding products