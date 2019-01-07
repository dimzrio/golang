package main

import "fmt"

func main() {

	// Example 1
	i := 1
	for i <= 5 {
		fmt.Println("Example1", i)
		i++
	}

	fmt.Println()

	// Example 2
	for i := 0; i < 5; i++ {
		fmt.Println("Example2", i)
	}
}
