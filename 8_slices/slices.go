package main

import "fmt"

func main() {
	// var nums []int
	// fmt.Println(nums == nil)

	//to avoid nil -> We can do so>>

	// var myNum = make([]int,2)
	// fmt.Println(myNum)

	 var mySlice = make([]int,5,6)
	//*Another way of doing the same 
	// var mySlice = []int {}
	// fmt.Println(mySlice)
	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 8)
	// fmt.Println(mySlice)
	// fmt.Println(cap(mySlice))
	var copySlice = make([]int , 4)
	copy(copySlice,mySlice)

	fmt.Println(mySlice,copySlice)


	//? slice operator 

	fmt.Println(mySlice[0:5]);
	fmt.Println(mySlice[0:])
}