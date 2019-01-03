package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("../basic_exampe/")
	data, err := box.FindString("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Contents of file:", data)

}
