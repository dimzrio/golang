package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "This is an example"
	var sha = sha1.New()
	sha.Write([]byte(text))
	fmt.Println(sha)

	var encrypt = sha.Sum(nil)
	fmt.Printf("Encrypted String: %x\n", encrypt)

}
