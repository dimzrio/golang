package main

import "fmt"

// CalculateCube func
func CalculateCube(number int, cubech chan int) {
	result := number * number * number
	cubech <- result
}

// CalculateSquare func
func CalculateSquare(number int, sqrch chan int) {
	result := number * number
	sqrch <- result
}

func main() {
	number := 10
	sqrch := make(chan int)
	cubech := make(chan int)

	go CalculateSquare(number, sqrch)
	go CalculateCube(number, cubech)

	resultSquare, resultCube := <-sqrch, <-cubech

	fmt.Println("Cube: ", resultCube)
	fmt.Println("Square: ", resultSquare)
}
