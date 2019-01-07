package main

import (
	"fmt"
)

func main() {

	// Alocate 2 space index 0,1
	var name [2]string
	name[0] = "Mo."
	name[1] = "Salah"
	fmt.Println("FullName: ", name[0], name[1])

	fmt.Println()

	// Declare 2 array with value
	var cat = [4]string{
		"Kuro",
		"Ema",
		"CM",
		"Nyiung",
	}
	fmt.Println("Cat No. 4:", cat[3])

	fmt.Println()

	// Declare array without set index
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Len:", len(numbers))

	fmt.Println()

	// Array multidimention
	var player = [3][2]string{ // 3 row, 2 coloumn
		[2]string{"Xherdan", "Shaqiri"},
		[2]string{"Mo.", "Salah"},
		[2]string{"Sadio", "Mane"},
	}
	fmt.Println("Player:", player)

	fmt.Println()

	// Sort declare
	programming := []string{"Golang", "Python"}
	fmt.Println("Programming:", programming)

	fmt.Println()

	// Initials array with struct
	cats := []struct {
		Name string
		Age  int
	}{
		{"Kuro", 4},
		{"Ema", 4},
		{"Centimeter", 3},
		{"Nyiung", 1},
	}

	fmt.Println(cats)

	fmt.Println()

	// Alocate array with make fucn
	makes := make([]string, 2)
	fmt.Println("Make func: ", makes)

}
