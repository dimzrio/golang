package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, errCreate := os.Create("example-bytes.txt")

	if errCreate != nil {
		log.Fatal(errCreate)
		return
	}

	content := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	line, errWrite := file.Write(content)

	if errWrite != nil {
		log.Fatal(errWrite)
		file.Close()
		return
	}

	fmt.Println(line, "bytes written successfully")

	errClose := file.Close()

	if errClose != nil {
		log.Fatal(errClose)
		return
	}

}
