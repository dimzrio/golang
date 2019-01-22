package main

import (
	"flag"
	"fmt"
)

func main() {

	// Example 1
	// var name = flag.String("name", "anonymous", "type your name")
	// var age = flag.Int64("age", 28, "type your age")

	// flag.Parse()

	// fmt.Println("Name:", *name)
	// fmt.Println("Age:", *age)

	// Example 2

	var name string
	flag.StringVar(&name, "name", "anonymous", "type your name")

	var age int64
	flag.Int64Var(&age, "age", 28, "type your age")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
