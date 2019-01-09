package main

import (
	"fmt"

	"github.com/dimzrio/golang-couses/03_packages/basic_package/strutils"
)

func main() {
	fmt.Println(strutils.GetInfo())
	fmt.Println("ReverseConcate:", strutils.ReverseConcate("Hello World"))
}
