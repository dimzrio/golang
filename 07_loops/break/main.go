package main

import "fmt"

func main() {

	// Example 1
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("Break")
			break
		} else {
			fmt.Println("Example", i)
		}
	}
}
