package main

import (
	"fmt"
	"time"
)

func Generator() chan string {
	ch := make(chan string)
	go func() {
		var input string
		for {
			fmt.Println("Input 10? ")
			fmt.Scan(&input)
			select {
			case ch <- input:
			case <-time.After(5 * time.Second):
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	n := <-number
	if n ==
	close(number)
}
