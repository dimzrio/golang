package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 123

	result := strconv.Itoa(num)
	fmt.Printf("Type %T: %s\n", result, result)
}
