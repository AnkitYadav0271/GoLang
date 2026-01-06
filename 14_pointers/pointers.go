package main

import "fmt"

//*passed by value
func changeNum(num int) {
	num = 5
	fmt.Println("ChangeNum num Memory :)", &num)
}

//* passed by reference

func changeNumReference(num *int) {
	*num = 8
	fmt.Println("At function:)", *num)
}

func main() {
	num := 1
	changeNumReference(&num)
	fmt.Println("At Main", num)
}
