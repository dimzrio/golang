package main

import (
	"fmt"
	"strings"
)

// HasPrefix is live a startwith in python
func main() {
	var startWith = strings.HasPrefix("This is an example string", "Thi")
	fmt.Println(startWith)
}
