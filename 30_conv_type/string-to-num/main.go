package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = "123"

	base := 8     // octa decimal
	typeInt := 64 // 32 or 64

	result, _ := strconv.ParseInt(num, base, typeInt)
	fmt.Printf("Type %T: %v\n", result, result)

	num = "1010"
	result, _ = strconv.ParseInt(num, 2, 8) // conv to string to binner with int8
	fmt.Printf("Type %T: %v\n", result, result)

}
