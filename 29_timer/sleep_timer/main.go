package main

import (
	"fmt"
	"time"
)

func sleep() {
	fmt.Println("start sleep")
	time.Sleep(time.Second * 4)
	fmt.Println("stop sleep after 4 seconds")
}

func main() {

	// Time Sleep
	sleep()
}
