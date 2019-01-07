package main

import "fmt"

func main() {

	point := 100

	// else if with temp variable
	if result := point / 3; result >= 50 {
		fmt.Println("Result greater than 50 :", result)
	} else if result >= 25 {
		fmt.Println("Result greater than 25 :", result)
	} else {
		fmt.Println("Result lower than 25 :", result)
	}
}
