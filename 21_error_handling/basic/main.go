package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Filename: %s", file.Name())
}
