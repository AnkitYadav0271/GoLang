package main

import (
	"fmt"
)

// func process(numChan chan int) {
// 	for {
// 		fmt.Println("Processing", <-numChan)
// 	}

// }

// func sum(result chan int, num1 int, num2 int) {
// 	res := num1 + num2
// 	result <- res
// }

// func task (done chan bool) {
// 	defer func(){done <- true}()
// 	fmt.Println("Processing...")
// }

func sendEmail(channel chan string, done chan bool) {
	func() {
		done <- true
	}()
	for email := range channel {
		fmt.Println("Email sending to ", email)
	}

}

func main() {
	//? Buffered Channel (Non blocking channel)
	channel := make(chan string, 10)
	done := make(chan bool)

	for i := 1; i <= 100; i++ {
		channel <- fmt.Sprintf("%d@gmail.com", i)
	}

	go sendEmail(channel, done)

	<-done

	channel <- "1@happy.com"
	channel <- "2@happy.com"
	channel <- "3@happy.com"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	//?UnBuffered channel (Blocking channel)
	// done := make(chan bool)
	// go task(done)
	// <- done
	// result := make(chan int)
	// go sum(result, 4, 6)
	// res := <-result
	// fmt.Println(res)
	// numChan := make(chan int)
	//*this is important send the data after function call other wise it will fall to deadlock condition
	// go process(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//  time.Sleep(time.Second * 2)
	//  messageChannel := make(chan string)

	// messageChannel <- "Hello Ping"

	// val := <-messageChannel

	// fmt.Println(val)
}
