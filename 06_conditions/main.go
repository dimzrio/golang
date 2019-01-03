package main

import "fmt"

func main() {
	x := 10
	y := 15

	// If else
	if x < y {
		fmt.Printf("x %d less then y %d\n", x, y)
	} else {
		fmt.Printf("x %d greter then y %d\n", x, y)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
