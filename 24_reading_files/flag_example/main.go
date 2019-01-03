package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	/***
	$ go build
	$ ./flag_example -ilepath=/Users/dimzrio/Documents/Projects/go/src/github.com/dimzrio/golang_course/24_reading_files/basic_exampe/example.txt
	***/
	path := flag.String("filepath", "text.txt", "file path to read from")
	flag.Parse()
	fmt.Println("Value files: ", *path)

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Content: ", string(data))
}
