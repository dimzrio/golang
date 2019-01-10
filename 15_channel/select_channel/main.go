package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func(channel1 chan<- string) {
		channel1 <- fmt.Sprintf("Channel 1 Stdout")
		time.Sleep(2 * time.Second)
	}(channel1)

	go func(channel2 chan<- string) {
		channel2 <- fmt.Sprintf("Channel 2 Stdout")
	}(channel2)

	select {
	case stdout1 := <-channel1:
		fmt.Println(stdout1)
	case stdout2 := <-channel2:
		fmt.Println(stdout2)
	}

}
