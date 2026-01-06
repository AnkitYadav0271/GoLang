package main

import "fmt"

//* In go we create enums with use of type and const because there is no enum keyword


type OrderStatus string

const (
	received OrderStatus = "received"                 
	confirmed = "confirmed"
	prepared = "prepared"
	delivered = "delivered"
)

func changeStatus(status OrderStatus) {
	fmt.Println("status changed to :",status)
}

func main() {
changeStatus(delivered)
}
