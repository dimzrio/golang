package main

import (
	"fmt"
	"os"
)

func main() {
	// bypass error
	files, _ := os.Open("example.txt")
	fmt.Println(files)
}
