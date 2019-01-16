package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	var url = "https://www.google.com"

	var encoded = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded:", encoded)

	var decoded, err = base64.URLEncoding.DecodeString(encoded)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Decode:", string(decoded))
}
