package main

import (
	"fmt"
	"math"

	"github.com/dimzrio/golang_course/03_packages/strutils"
)

func main() {
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutils.ReverseConcate("Hello World - ReverseConcate"))
	fmt.Println(strutils.ReverseBuffer("Hello World - ReverseBuffer"))
	fmt.Println(strutils.ReverseJoin("Hello World - ReverseJoin"))
}
