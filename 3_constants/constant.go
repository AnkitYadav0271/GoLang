package main

import "fmt"
const SurName = "Yadav"
func main() {
	const Name string = "Happy"
	fmt.Println(Name)
	fmt.Println(SurName)

	const (port = 6000
	host= "localhost")

	fmt.Println(port)
	fmt.Println(host)
}