package main

import "fmt"

func main() {

	// Anonymouse struct is struct not declare in first code, but if needed
	player1 := struct {
		Name, Club string
		Age        int
	}{
		Name: "Mo. Salah",
		Club: "Liverpool",
		Age:  28,
	}

	fmt.Println(player1)
}
