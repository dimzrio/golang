package main

import (
	"fmt"
	"os"
)

func main() {

	// Assertion
	file, err := os.Open("example.txt")

	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}

	fmt.Println(file.Name(), "Successfully to open")
}
