package main

import (
	"fmt"
)

func main() {
	var cat = [4]string{
		"Kuro",
		"Ema",
		"CM",
		"Nyiung",
	}

	for index, name := range cat {
		fmt.Printf("%d. %s\n", index, name)
	}
}
