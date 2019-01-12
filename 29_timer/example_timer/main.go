package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcherTimeout(ch <-chan bool) {
	if <-ch {
		fmt.Printf("\nTimeout, exit 0\n")
		os.Exit(0)
	}
}

func main() {
	timeout := 5
	ch := make(chan bool)

	var input string

	go timer(timeout, ch)
	go watcherTimeout(ch)

	for i := 1; i <= 3; i++ {

		fmt.Print("[?] What is 7 + 3? ")
		fmt.Scan(&input)

		if input == "10" {
			fmt.Println("Correct")
			break
		} else {
			fmt.Println("Incorrect")
		}
	}
}
