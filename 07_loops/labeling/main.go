package main

import "fmt"

func main() {

	// Label example loops
example1:
	for i := 1; i <= 10; i++ {

		for j := 1; j <= 3; j++ {

			if i == 5 {
				break example1
			}

			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}
