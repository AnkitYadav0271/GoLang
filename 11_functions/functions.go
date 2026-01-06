package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func languages() (string, string, string) {
	return "golang", "typescript", "c"
}

// func processIt(fn func(a int) int) {
// 	fn(4)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 2 * a
	}
}

func main() {
	// result := add(5, 6)
	// fmt.Println("Result is here", result)

	// lang1, lang2, lang3 := languages()
	// fmt.Println(lang1, lang2, lang3)

	// fn := func(a int) int {
	// 	fmt.Println(2 * a)
	// 	return 2 * 2
	// }

	fn := processIt()
	fmt.Println((fn(7)))
}
