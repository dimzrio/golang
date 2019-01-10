package main

import (
	"fmt"
	"sync"
	"time"
)

func getInt(i int, result *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	result.Done()
}

func main() {

	no := 5
	var result sync.WaitGroup

	for i := 0; i < no; i++ {
		result.Add(1)
		go getInt(i, &result)
	}

	result.Wait()
	fmt.Println("All go routines finished executing")
}
