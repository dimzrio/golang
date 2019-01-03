package main

import (
	"fmt"
	"time"
)

func getInt(result chan int) {

	for i := 0; i < 10; i++ {
		result <- i
		fmt.Println("Success strore i: ", i, " to channel")
	}
	close(result)
}

func main() {
	// Set channel buffer with 3 capacity for unblocking
	// blocking channel if capacity is full
	result := make(chan int, 3)

	// uncommant below for compare without capacity
	// result := make(chan int)
	go getInt(result)
	time.Sleep(1 * time.Second)

	for i := range result {
		fmt.Println("Read value ", i, " from channel")
		time.Sleep(1 * time.Second)
	}

}
