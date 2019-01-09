package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // alocate 2 core cpu for run this program

	// task 1
	go func() {
		for i := 1; i <= 20; i++ {
			fmt.Println("Task-1 ->", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// task 2
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Task-2 ->", i)
			time.Sleep(4 * time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input) // trap all stdout before user press "Enter" (Blocking)

	fmt.Println("Done")
}
