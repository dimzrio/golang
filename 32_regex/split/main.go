package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "This,is,an,example,regex,of,golang"
	regex, _ := regexp.Compile(`,+`)
	isMatch := regex.Split(text, -1)

	for _, value := range isMatch {
		fmt.Println(value)
	}
}
