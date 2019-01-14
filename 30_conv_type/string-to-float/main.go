package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "24.12"
	var result, err = strconv.ParseFloat(str, 32)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type %T: %f\n", result, result)
}
