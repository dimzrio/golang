package main

import (
	"fmt"
	"strings"
)

func main() {
	var countT = strings.Count("This is an example string", "t")
	fmt.Println(countT)
}
