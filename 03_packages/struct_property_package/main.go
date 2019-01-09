package main

import (
	"fmt"

	. "github.com/dimzrio/golang-couses/03_packages/struct_property_package/authors" // dot bypass package name
)

func main() {

	// Example bypass
	setAuthor := AuthorInfo{Name: "Dimas", Email: "dimas@gmail.com"}
	fmt.Println(setAuthor.GetAuthor())

}
