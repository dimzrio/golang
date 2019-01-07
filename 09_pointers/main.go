package main

import "fmt"

func main() {

	// Pointer is assign variable to memory for increase performance
	a := 5
	b := &a // & is assign pointer to var a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Type a : %T\n", a)
	fmt.Printf("Type b : %T", b)

	// call value of b with 2 ways
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value ponter b
	*b = 10 // * representing pointer
	fmt.Println(a)

}
