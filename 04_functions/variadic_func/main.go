package main

import "fmt"

// Variadic args func in golang like *args in python
func addition(numbers ...int) int {
	result := 0

	for _, val := range numbers {
		result += val
	}

	return result
}

func main() {
	add := addition(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Result:", add)
}
