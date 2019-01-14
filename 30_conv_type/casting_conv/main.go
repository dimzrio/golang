package main

import "fmt"

func main() {

	// float to int
	float := 24.000
	resultFloat := int64(float)
	fmt.Println(resultFloat)

	// float to int
	resultInt := int64(123.00)
	fmt.Println(resultInt)

	// string to byte
	text := "Dimas"
	data := []byte(text)
	for _, b := range data {
		fmt.Printf("Char %s -> Byte %v\n", string(b), b)
	}
}
