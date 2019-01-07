package main

import "fmt"

func strfunc(names []string, f func(string) string) []string {
	var result []string

	for _, n := range names {
		// Run func args with value name
		result = append(result, f(n))
	}

	return result
}

func main() {
	names := []string{"example1", "example2", "example3"}

	// Send func as args
	result := strfunc(names, func(n string) string {
		return "Hello " + n
	})

	for index, nm := range result {
		fmt.Printf("%d. %s\n", index, nm)
	}
}
