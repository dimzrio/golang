package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	var text = "This is an example of regex"
	regex, err := regexp.Compile(`[a-z]+`) // find char one or more a - z

	if err != nil {
		log.Println(err)
	}

	isMatch := regex.MatchString(text)

	fmt.Println(isMatch) // true -> because regex is metch in string
}
