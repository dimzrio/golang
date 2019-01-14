package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	var newText4 = strings.Replace(text, "banana", "Choco Banana", -1)
	fmt.Println(newText4)

	// Replace with NewReplace
	textRep := "This is an example replace"
	replace := strings.NewReplacer(
		"This", "Ini",
		"is", "adalah",
		"an", "sebuah",
		"example", "contoh",
		"replace", "mengganti",
	)
	textRep = replace.Replace(textRep)
	fmt.Println(textRep)

}
