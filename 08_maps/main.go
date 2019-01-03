package main

import "fmt"

func main() {

	// Declare map
	emails := make(map[string]string)

	// Assign
	emails["rio"] = "rio@gmail.com"
	emails["hana"] = "hana@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["rio"])

	// Shortened
	family := map[string]int{"rio": 28, "hana": 28, "asiyah": 1, "nyiung": 1}
	fmt.Println(family)
	fmt.Println(len(family))

	// Delete
	delete(family, "nyiung")
	fmt.Println(family)

}
