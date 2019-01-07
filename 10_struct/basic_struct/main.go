package main

import "fmt"

// Player structure
type Player struct {
	Name, Club string // First char must be Capitalize
	Age        int
}

func main() {

	p1 := Player{
		Name: "Mo. Salah",
		Club: "Liverpool",
		Age:  28,
	}

	fmt.Println(p1)
}
