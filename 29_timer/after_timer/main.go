package main

import (
	"fmt"
	"time"
)

func main() {

	// time.After like a time.Sleep but time.After return to channel
	fmt.Println("start after time")
	<-time.After(2 * time.Second)
	fmt.Println("stop after time after 2 second")
}
