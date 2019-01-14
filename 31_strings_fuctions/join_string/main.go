package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"This", "is", "an", "example"}

	str := strings.Join(data, "-")
	fmt.Println(str)
}
