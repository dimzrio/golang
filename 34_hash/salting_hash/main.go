package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func saltHash(str string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltText = fmt.Sprintf("%s%s", str, salt)

	sha := sha1.New()
	sha.Write([]byte(saltText))
	result := fmt.Sprintf("%x", sha.Sum(nil))

	return result, salt
}

func main() {
	var text = "This is an example"
	encrypted, salt := saltHash(text)

	fmt.Printf(`
Text : %s Salt: %s
Encrypted : %s
`, text, salt, encrypted)

	fmt.Printf(`
Text : %s Salt: %s
Encrypted : %s
`, text, salt, encrypted)

}
