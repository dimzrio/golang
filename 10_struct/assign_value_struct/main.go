package main

import "fmt"

// Player structure
type Player struct {
	Name, Club string // First char must be Capitalize
	Age        int
}

func main() {
	// assign value
	var p1 Player
	p1.Name = "Mo. Salah"
	p1.Club = "Liverpool"
	p1.Age = 28

	fmt.Println(p1)
}
