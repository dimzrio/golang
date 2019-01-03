package main

import (
	"fmt"
	"time"
)

func procedur(channel chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		channel <- i
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go procedur(channel)
	for i := range channel {
		fmt.Println("Receiver: ", i)
	}
}
