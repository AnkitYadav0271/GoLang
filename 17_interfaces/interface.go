package main

import "fmt"

type razorpay struct{}

//* Here comes use of interface

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (r razorpay) pay(amount float32) {
	fmt.Println("paying through razorpay.....:",amount)
}

type stripe struct {}

func (s stripe) pay(amount float32) {
	fmt.Println("Paying through stripe amount of :",amount)
}

func main() {
// razorpayPaymentGw := razorpay{}
stripePaymentGw := stripe{}
paymentGw := payment{
	gateway: stripePaymentGw,
}
paymentGw.gateway.pay(100)
}