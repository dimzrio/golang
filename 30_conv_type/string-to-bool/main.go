package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "true"
	var result, err = strconv.ParseBool(str)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type %T: %v\n", result, result)
}
