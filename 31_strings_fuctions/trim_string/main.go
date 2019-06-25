package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\t    Hello, World\n    "
	t := strings.TrimSpace(s)
	fmt.Printf("%s\n", t)
}
