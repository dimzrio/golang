package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	result, error := division(4, 0)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("%.2f\n", result)

}

// division func return fload and error value
func division(x, y float64) (float64, error) {

	if y == 0 || reflect.TypeOf(y).Kind() == reflect.String {

		// Call error packages with fuction new
		return 0, errors.New("Division must be grether then 0 and is not string")
	}

	result := x / y
	return result, nil
}
