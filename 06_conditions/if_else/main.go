package main

import "fmt"

func main() {
	x := 10
	y := 15

	if x < y {
		fmt.Printf("x %d less then y %d\n", x, y)
	} else {
		fmt.Printf("x %d greter then y %d\n", x, y)
	}
}
