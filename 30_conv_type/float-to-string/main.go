package main

import (
	"fmt"
	"strconv"
)

func main() {
	var float = 24.12
	var result = strconv.FormatFloat(float, 'f', 3, 64) // f : enable format exponent, with 3 digit type float64

	fmt.Printf("Type %T: %v\n", result, result)
}
