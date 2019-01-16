package main

import (
	"fmt"
	"regexp"
)

// https://github.com/google/re2/wiki/Syntax
func main() {

	var text = "This is an example of regex"
	regex, _ := regexp.Compile(`[a-z]+`) // compile will be change to regex.*Regexp type

	isMatch := regex.ReplaceAllString(text, "T") // TT T T T T T
	fmt.Println(isMatch)
}
