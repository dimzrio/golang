package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	var text = "This is an example of regex"

	regex, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		log.Println(err)
	}

	isMatch := regex.FindStringIndex(text)

	startIndex := isMatch[0]
	endIndex := isMatch[1]

	str := text[startIndex:endIndex]

	fmt.Println(str) // his
}
