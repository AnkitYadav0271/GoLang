package main

import (
	"fmt"
	"maps"
)

//Maps -> hash ,object,set in js

func main() {
	//creating a map 
	m:= make(map[string]string)
	//setting element 
	m["name"] = "Happy"
	fmt.Println("Here is the output: ",m["age"])

	m["address"] = "Raipur Bharkhi Baswahi kunda pratapgarh"

	delete(m,"name")

	// fmt.Println(m)

	clear(m)
	// fmt.Println(m)

	myMap := map[string] int {"rollNo":20,"age":19}

	v,ok := myMap["age"]

	fmt.Println(v)

	if ok {
		fmt.Println("Ok",ok)
		
	}else {
		fmt.Println("Not ok")
	}

	// fmt.Println(myMap)

	myMap1 := map[string] int {"rollNo":20,"age":19}
	myMap2 := map[string] int {"rollNo":20,"age":19}

	fmt.Println(maps.Equal(myMap1,myMap2))
}