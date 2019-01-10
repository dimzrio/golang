package main

import "fmt"

// fuction addMumber with args int array and
func addNumber(a []int, add10 func(number int) int) []int {
	var result []int
	for _, value := range a {
		result = append(result, add10(value))
	}
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// send data to addNumber func with args int of array & fuction add 10
	result := addNumber(data, func(value int) int {
		return value + 10
	})
	fmt.Println(result)
}
