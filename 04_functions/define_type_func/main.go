package main

import "fmt"

type div func(x, y float64) float64

func main() {
	var d div = func(x, y float64) float64 {
		return x / y
	}

	result := d(7, 5)
	fmt.Printf("Division: %.2f\n", result)
}
