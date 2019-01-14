package main

import (
	"fmt"
	"strings"
)

// hasSuffix to check endwith string
func main() {
	var endWith = strings.HasSuffix("This is an example string", "ng")
	fmt.Println(endWith)
}
