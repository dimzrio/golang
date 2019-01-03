package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatal(err)
		return
	}

	content := "Go Programming"

	_, errWrite := fmt.Fprintln(file, content)

	if errWrite != nil {
		log.Fatal(errWrite)
		file.Close()
		return
	}

	errClose := file.Close()
	if errClose != nil {
		log.Fatal(errClose)
		return
	}

	fmt.Println("Append successfully")

}
