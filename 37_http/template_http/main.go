package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(resp http.ResponseWriter, req *http.Request) {
	path := req.RequestURI

	respWriter := fmt.Sprintf("Path: %s \nThis is homepage", path)
	fmt.Fprintln(resp, respWriter)
}

func about(resp http.ResponseWriter, req *http.Request) {
	var data = map[string]string{
		"name": "Dimas",
		"text": "This is an example",
	}

	t, err := template.ParseFiles("about.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(resp, data)

}
func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello World")
	})

	http.HandleFunc("/home", home)

	http.HandleFunc("/about", about)

	fmt.Println("Starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
