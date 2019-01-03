package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

/***
$ go build
$ ./read_line_by_lines -filepath=/Users/dimzrio/Documents/Projects/go/src/github.com/dimzrio/golang_course/24_reading_files/basic_exampe/example.txt
***/

func main() {
	path := flag.String("filepath", "text.txt", "file path to read from")
	flag.Parse()

	data, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	content := bufio.NewScanner(data)

	for content.Scan() {
		fmt.Println(content.Text())
		time.Sleep(1 * time.Second)
	}

	errScan := content.Err()

	if errScan != nil {
		log.Fatal(errScan)
	}

	defer func() {
		if err = data.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}
