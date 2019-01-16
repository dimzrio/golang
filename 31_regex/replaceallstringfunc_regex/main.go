package main

import (
	"fmt"
	"regexp"
)

// https://github.com/google/re2/wiki/Syntax
func main() {

	var text = "This is an example of regex"
	regex, _ := regexp.Compile(`[a-z]+`) // compile will be change to regex.*Regexp type

	isMatch := regex.ReplaceAllStringFunc(text, func(m string) string {
		if m == "regex" {
			return "regular expresion"
		}
		return m
	})
	fmt.Println(isMatch)
}
