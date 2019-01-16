package main

import (
	"fmt"
	"log"
	"regexp"
)

// https://github.com/google/re2/wiki/Syntax
func main() {
	regex, err := regexp.Compile(`[a-z]+`) // compile will be change to regex.*Regexp type

	if err != err {
		log.Panicln(err)
	}

	fmt.Printf("Type: %T\n", regex)
}
