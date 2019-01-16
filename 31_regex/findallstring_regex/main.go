package main

import (
	"fmt"
	"regexp"
)

// https://github.com/google/re2/wiki/Syntax
func main() {

	var text = "This is an example of regex"
	regex, _ := regexp.Compile(`[a-z]+`) // compile will be change to regex.*Regexp type

	isMatch := regex.FindAllString(text, -1) // -1 is mean get all array
	fmt.Println(isMatch)

	isMatch = regex.FindAllString(text, 1) // 0 get array index 0
	fmt.Println(isMatch)
}
