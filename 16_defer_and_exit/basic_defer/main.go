package main

import "fmt"

func main() {
	defer fmt.Println("Di Eksekusi terakhir 2")
	defer fmt.Println("Di Eksekusi terakhir 1")
	fmt.Println("Di Eksekusi pertama")

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i) // with defer stdout reverse
	}

	// os.exit(0) --> if using os.exit, defer not execute
}
