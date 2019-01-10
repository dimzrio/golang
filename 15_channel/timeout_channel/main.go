package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(data chan<- int) {
	// while true
	for i := 0; true; i++ {
		data <- i
		sec := time.Duration(rand.Intn(10)) * time.Second
		fmt.Println("Timeout: ", sec)
		time.Sleep(sec) // Sleep Random 0 - 5
	}
}

func main() {
	data := make(chan int)
	go sendData(data)

loop:
	for {
		select {
		case id := <-data:
			fmt.Println("Task->", id)
		case <-time.After(5 * time.Second): // timeout 5 second
			fmt.Println("Break - Timeout more than 5 second")
			break loop
		}
	}
}
