package main

import "fmt"

func main() {
	// Array data with map value
	data := []map[string]string{
		map[string]string{"player": "Mo.Salah", "club": "Liverpool"},
		map[string]string{"player": "Sadio Mane", "club": "Liverpool"},
		map[string]string{"player": "Lewandoski", "club": "As Roma"},
	}

	for id, key := range data {
		fmt.Printf("%d. %s -> %s\n", id, key["player"], key["club"])
	}
}
