package main

import (
	"fmt"

	"github.com/dimzrio/golang_course/18_oop/family"
)

func main() {
	fmt.Println("Classes - OOP in Golang")
	/*
		data := family.Family{
			Name:    "Dimas Rio",
			Gander:  "L",
			Age:     28,
			Merried: true,
		}
	*/

	data := family.New("Dimas Rio", "L", 28, true)

	fmt.Println(data.Info())
}
