package main

import "fmt"

//? Variadic functions are rest operator in js

func sum(nums ...int) int {
	total := 0
	for num := range nums {
		total += num
	}
	fmt.Println(nums)
	return total
}

func main() {
	fmt.Println("hello main")
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 8, 8, 9))
}
