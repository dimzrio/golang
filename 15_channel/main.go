package main

import "fmt"

// Sends and receives to a channel are blocking by default
// data := <- a          >>>>> read from channel a
// a <- data             >>>>> write to channel a

/*
Note:
Jika di dalam func hello tidak ada komunikasi dengan channel lain
maka cukup menggunakan result <- true

tapi jika ada value yg di write ke channel lain maka cukup di close saja
ex:

var task = make(chan int)
func createTaks(noTaks int) {
	for i := 1; i <= noTaks; i++ {
		task <- i
	}
	close(task)
}

func main() {
	noTaks := 10

	go createTaks(noTaks)

	for taskValue := range task {
		fmt.Println(taskValue)
	}
}
*/

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
