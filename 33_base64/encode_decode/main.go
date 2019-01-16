package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var text = "This is an example of golang"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	fmt.Println(encoded) // [ 0 0 0...0 0 0 ]
	base64.StdEncoding.Encode(encoded, []byte(text))
	fmt.Println("Encoded:", encoded) // [86 71 104...119 61 61]

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	fmt.Println(decoded) // [ 0 0 0...0 0 0 ]
	base64.StdEncoding.Decode(decoded, encoded)
	fmt.Println("Decoded:", string(decoded))

}
