package main

import "fmt"

func main() {
	// Interface as type data, you can assign any value
	var data interface{}

	// Set as float64
	data = 3.14
	fmt.Println(data)

	data = []string{"example1", "example2"}
	fmt.Println(data)

	data = map[int]string{1: "example1", 2: "example2"}
	fmt.Println(data)

	// Assing interface to map for more flexibility
	dict := make(map[int]interface{})

	dict[1] = map[string]interface{}{"name": "Mario Mandzukic", "club": "Juventus", "age": 32}
	dict[2] = map[string]interface{}{"name": "E. Dzeko", "club": "AS Roma", "age": 32}
	dict[3] = []string{"AS Roma", "Juventus", "Intermilan"}
	dict[4] = "End of contract 2019"

	fmt.Println(dict)
}
