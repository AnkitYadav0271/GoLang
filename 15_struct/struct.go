package main

import "fmt"

type order struct {
	id     string
	amount float32
	status string
}
func (o *order) changeStatus (status string) {
	o.status = status
}

func newOrder (id string , amount float32 , status string) *order {
	myOrder:= order{
		id: id,
		amount: amount,
		status: status,
	}

	return & myOrder
}

func (o * order) showAmount () float32 {
	return o.amount
}




func main() {
	// myOrder := order{
	// 	id:     "12",
	// 	amount: 50.5,
	// 	status: "pending",
	// }
    // myOrder.amount= 56
	// fmt.Println(myOrder)
	// myOrder.changeStatus("hello")
	// fmt.Println(myOrder.showAmount())
	// fmt.Println(myOrder)

	// againOder:=newOrder("2",56.4,"pending")
	// fmt.Println(againOder)
	language:= struct {
	name string
	isGood bool
}{
"golang",
true,
}

	fmt.Println(language);
}
