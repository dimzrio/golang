package main

import "fmt"

func main() {

	// Declare map
	emails := make(map[string]string)
	emails["rio"] = "rio@gmail.com"
	emails["hana"] = "hana@gmail.com"
	fmt.Println(emails)
	fmt.Println("Len: ", len(emails))
	fmt.Println(emails["rio"])

	// Shortened
	family := map[string]int{"rio": 28, "hana": 28, "asiyah": 1, "nyiung": 1}
	delete(family, "nyiung") // delete key nyiung
	fmt.Println(family)
}
