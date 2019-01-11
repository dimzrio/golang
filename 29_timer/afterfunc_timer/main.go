package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	fmt.Println("Start afterfunc")
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Afterfunc Callback")
		ch <- true
	})
	<-ch
	fmt.Println("Stop afterfuct after 2 second")
}
