package main

import "fmt"

func main() {

	// ref: https://www.tutorialspoint.com/go/go_data_types.htm
	var f float64 = 5.4
	var f1, f2 float64 = 5.4, 5.4
	f2d := 5.4787

	fmt.Printf("Type %T: %f\n", f, f)
	fmt.Printf("Type %T: %.2f\n", f2d, f2d)
	fmt.Println("Float: ", f1, ",", f2)
}
