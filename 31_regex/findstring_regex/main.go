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

	isMatch := regex.FindString(text)

	fmt.Println(isMatch) // his -> because regex is not match with T
}
