package main

import "fmt"

// Sends and receives to a channel are blocking by default
// data := <- a          -> read from channel a and assign to var data
// a <- data             -> write value data to channel a

func hello(result chan bool) {
	fmt.Println("Hello world goroutine")

	// write to result
	result <- true
}

func main() {
	// Defince channel
	result := make(chan bool)

	// run goroutine
	go hello(result)

	// read value from channel but not used
	// <- result

	// read value from channel
	resultChannel := <-result

	fmt.Println("result channel: ", resultChannel)
	fmt.Println("main function")

}
