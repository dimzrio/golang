package main

import "fmt"

func main() {
	// Anonymous func
	result := func() string {
		return fmt.Sprintf("This is an example")
	}
	fmt.Println(result())

	// Anonymous func with args
	resultArgs := func(args string) string {
		return fmt.Sprintf("Hello %s", args)
	}
	fmt.Println(resultArgs("Dimas"))
}
