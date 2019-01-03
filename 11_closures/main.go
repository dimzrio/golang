package main

import "fmt"

func main() {

	// Closures like a lambda in python
	sum := func(n int) int {
		result := 0
		for i := 0; i < n; i++ {
			result += i
		}
		return result
	}

	fmt.Println(sum(15))
}
