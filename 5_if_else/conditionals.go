package main

import "fmt"

func main() {
	age := 17
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is Minor")
	}

	var role = "admin"
	var hasPermission = false
	if role == "admin" || hasPermission {
		fmt.Println("Grant")
		fmt.Println("At or operator")
	}

	if role == "admin" && hasPermission {
		fmt.Println("grant")
		fmt.Println("At and operator")
	}

	if old := 16; old >= 18 {
		fmt.Println("Matured")
	} else if age > 12 {
		fmt.Println("Teen")
	}
}
