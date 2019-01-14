package main

import (
	"fmt"
	"strconv"
)

func main() {
	var boolean = true
	var result = strconv.FormatBool(boolean)

	fmt.Printf("Type %T: %s\n", result, result)
}
