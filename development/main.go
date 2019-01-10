package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	name, _ := input.ReadString('\n')

	fmt.Println(name)
}
