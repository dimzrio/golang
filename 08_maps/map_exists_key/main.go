package main

import "fmt"

func main() {

	// Declare map
	data := map[int]string{1: "Mo. Salah", 2: "Sadio Mane"}

	// Check key in map
	_, isExist := data[3] // return boolean
	if isExist {
		fmt.Println(isExist)
	} else {
		fmt.Println("Key doesn't exist")
	}
}
