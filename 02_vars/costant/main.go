package main

import "fmt"

func main() {
	// constanta is a fix variable and can't assign
	const phi float64 = 3.14

	//phi = 3.15

	fmt.Printf("Type %T: %f\n", phi, phi)

}
