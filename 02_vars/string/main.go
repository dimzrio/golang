package main

import "fmt"

func main() {

	// ref: https://www.tutorialspoint.com/go/go_data_types.htm
	var name string = "dimzrio"
	var firstname, lastname string = "dimas", "rio"
	names := "dimzrio"

	fmt.Printf("Type %T: %s\n", name, name)
	fmt.Printf("Type %T: %s\n", names, names)
	fmt.Println("Name:", firstname, lastname)
}
