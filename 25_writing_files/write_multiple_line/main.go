package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, errCreate := os.Create("example.txt")

	if errCreate != nil {
		log.Fatal(errCreate)
		return
	}

	content := []string{"This is an example", "Write multiple lines", "Golang Programming language"}

	for _, writes := range content {

		_, errWrite := fmt.Fprintln(file, writes)

		if errWrite != nil {
			log.Fatal(errWrite)
			file.Close()
			return
		}
	}

	errClose := file.Close()

	if errClose != nil {
		log.Fatal(errClose)
		return
	}

	fmt.Println("Writeln successfully")
}
