package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {

	newOrder := order{
		id:        "1",
		amount:    30,
		status:    "received",
		createdAt: time.Now(),
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	newOrder.customer.name = "Saidur"
	fmt.Println(newOrder)
}
