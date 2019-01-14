package main

import (
	"fmt"
	"strings"
)

func main() {
	var getIndex = strings.Index("This is an example string", "is")
	fmt.Println(getIndex) // get result index 2, T -> 0, h -> 1, i(s) -> 2
}
