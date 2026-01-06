package main

import (
	"fmt"
	"time"
)
type order struct {
	id string
	price float32
	delivered bool
	time time.Time
	customer
}

type customer struct {
	name string
	address string
	contact string
	email string
}

func main () {
newCustomer :=customer {
name: "Happy",
contact:"7800958614",
address: "Raipur Bharkhi Baswahi Kunda Pratapgarh 230204",
email: "ankitofficial0271@gmail.com",
}

newOrder := order {
	id: "1",
	price: 34.45,
	delivered: true,
	time: time.Now(),
	customer: newCustomer,
}

fmt.Println(newOrder)
}