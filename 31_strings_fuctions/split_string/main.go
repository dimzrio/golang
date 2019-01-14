package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This is an example string"
	var data1 = strings.Split(text, " ")
	for index, str := range data1 {
		fmt.Println(index, str)
	}

	var data2 = strings.Split(text, "")
	fmt.Println(data2)
}
