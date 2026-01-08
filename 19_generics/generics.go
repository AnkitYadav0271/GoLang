package main

import (
	"fmt"
)

func printSlice[T string | int](items []T) {
	for _, value := range items {
		fmt.Println(value)
	}
}

type Stack [T any] struct {
	element []T
}

func main() {
	myStack := Stack [int] {
		element: []int {123,123,123},
	}

	fmt.Println(myStack.element);
// val:=[]string {"Hello","Happy"}
// printSlice(val);
}
