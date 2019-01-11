package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("Time now: %v\n", t)
	fmt.Println("year:", t.Year(), "month:", t.Month())

	t2 := time.Date(2019, 01, 11, 07, 1, 0, 0, time.UTC)
	fmt.Printf("Set time: %v\n", t2)

}
