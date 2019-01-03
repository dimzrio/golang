package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	line, err := file.WriteString("This is an example write file")

	if err != nil {
		log.Fatal(err)
		file.Close()
		return
	}

	fmt.Println(line, "Bytes write successfully...")

	errClode := file.Close()

	if errClode != nil {
		log.Fatal(errClode)
	}

}
