package main

import (
	"fmt"
	"time"
)

func hello(result chan bool) {
	fmt.Println("Hello goroutine going to sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine awake to sleep")
	result <- true
}

func main() {
	fmt.Println("Starting gorouting..")
	result := make(chan bool)
	go hello(result)
	<-result
	fmt.Println("Done")
}
