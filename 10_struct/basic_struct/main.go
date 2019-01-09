package main

import "fmt"

// Player structure
type Player struct {
	Name, Club string // First char must be Capitalize for Public, accessable from another package
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
