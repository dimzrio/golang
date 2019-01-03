package main

import "fmt"

func main() {

	// ref: https://www.tutorialspoint.com/go/go_data_types.htm

	// String
	var name = "rio"
	var email = "rio@gmail.com"

	// int
	var age int8 = 28

	// Boolean
	var isCool = true

	// Float
	var size = 1.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("Type isCool %T size: %T\n", isCool, size)

	fmt.Println()

	// Shortened
	name, age, email, isCool, size = "hana", 28, "hana@gmail.com", true, 1.5
	fmt.Println(name, age, isCool, email)
	fmt.Printf("Type isCool %T size: %T\n", isCool, size)

	// List
	type lists [2]string
	data := lists{"Hallo", "World"}
	fmt.Println(data)
}
