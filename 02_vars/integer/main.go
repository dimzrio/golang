package main

import "fmt"

func main() {

	// ref: https://www.tutorialspoint.com/go/go_data_types.htm
	var age int8 = 28
	var age1, age2 int = 28, 28
	ages := 28

	fmt.Printf("Type %T: %d\n", age, age)
	fmt.Printf("Type %T: %d\n", ages, ages)
	fmt.Println("Age:", age1, ",", age2)
}
