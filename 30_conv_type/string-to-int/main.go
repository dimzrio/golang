package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = "123"

	result, err := strconv.Atoi(num)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type %T: %d\n", result, result)
}
