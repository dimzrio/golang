package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	var text = "This is an example"

	regex, err := regexp.Compile(text)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(regex)
}
