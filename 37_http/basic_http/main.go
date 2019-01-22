package main

import (
	"fmt"
	"net/http"
)

func home(resp http.ResponseWriter, req *http.Request) {
	path := req.RequestURI

	respWriter := fmt.Sprintf("Path: %s \nThis is homepage", path)
	fmt.Fprintln(resp, respWriter)
}

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello World")
	})

	http.HandleFunc("/home", home)

	fmt.Println("Starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
