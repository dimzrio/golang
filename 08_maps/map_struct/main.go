package main

import "fmt"

func main() {

	// Map with struct
	type Employee struct {
		Name    string
		Age     int
		Merried bool
	}
	data := map[int]Employee{
		1: {Name: "Rio", Age: 28, Merried: true},
		2: {Name: "Dimas", Age: 28, Merried: true},
	}
	fmt.Println(data)
	fmt.Println(data[2])
}
