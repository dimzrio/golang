package main

import "fmt"

func main() {
	// Array
	var CatArray [2]string

	// Assign
	CatArray[0] = "Kuro"
	CatArray[1] = "Ema"

	fmt.Println(CatArray)

	// Shortened
	MyCats := []string{"Kuro", "Ema", "Nyiung", "Centimeters"}
	fmt.Println(MyCats)

	// Slices
	fmt.Println(MyCats[1:3]) // start at index 1 stop at index 3

}
