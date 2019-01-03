package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// With index
	for index, id := range ids {
		fmt.Printf("%d - ID: %d\n", index, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Sum all value
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d", sum)

	// Maps
	ages := map[string]int{"rio": 28, "hana": 28, "asiyah": 1}
	for name, age := range ages {
		fmt.Printf("My name is %s, I'm %d years old\n", name, age)
	}
}
