package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 1
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	{
	// 		fmt.Println("other")
	// 	}
	//}

	//? Multiple conditional switch case

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")

	default:
		fmt.Println("working day")
	}

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its integer")
		case string:
			fmt.Println("Its string")
		case bool:
			fmt.Println("Its boolean")
		default:
			fmt.Println("Other")

		}
	}

	whoAmI(true)

}
