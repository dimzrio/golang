package main

import (
	"fmt"
	"strings"
)

func main() {
	// Array
	var CatArray [2]string

	// Assign
	CatArray[0] = "Kuro"
	CatArray[1] = "Ema"

	fmt.Println(CatArray)

	// Shortened
	MyCats := []string{"Kuro", "Ema", "Nyiung", "Centimeters"}

	// Append lists
	MyCats = append(MyCats, "Emy")
	fmt.Println(MyCats)

	// Slices
	fmt.Println(MyCats[1:3]) // start at index 1 stop at index 3

	// Initials array with struct
	Cats := []struct {
		Name string
		Age  int
	}{
		{"Kuro", 4},
		{"Ema", 4},
		{"Centimeter", 3},
		{"Nyiung", 1},
	}

	fmt.Println(Cats)

	// Slice Capacity & Length
	fmt.Println("Cap: ", cap(MyCats))
	fmt.Println("Len: ", len(MyCats))

	// Allocation Array with make fuctions
	makes := make([]string, 0, 5)
	fmt.Println(cap(makes))

	s := makes[:2]
	fmt.Println(len(s))

	// Multiple slice array
	board := [][]string{ // --> [] [] cz only 2 col & row
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	fmt.Println(board)

	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
