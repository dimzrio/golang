package main

import (
	"fmt"
	"strings"
)

func splitName(name string) (string, string) {
	str := strings.Split(name, " ")
	return str[0], str[1]
}

func main() {
	firstName, lastName := splitName("Mario Tosca")
	fmt.Println("Firstname:", firstName)
	fmt.Println("Lastname:", lastName)
}
