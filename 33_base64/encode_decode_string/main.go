package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	var text = "This is an example of golang"

	// Encode
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Encode:", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		log.Fatal(err)
		return
	}

	result := string(decoded)
	fmt.Println("Decode:", result)
}
