package main

import "fmt"

//only construct in go for looping
func main () {
	//there is no while loop in go 
	//so we have to make it by own
	i:=1
	for  i <= 3 {
		fmt.Println(i)
		i++
	}

	//* classic for loop 

	for i:= 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for i:= range 11 {
		fmt.Println(i)
	}
	
}