package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("This is an example string", "example")
	fmt.Println(isExists)
}
