package main

import (
	"fmt"
	"os"
)

func main() {
	var getArgs = os.Args
	fmt.Println("Get all Args:", getArgs[1:])
}
