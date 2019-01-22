package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://example.com/visual/2018/10/26/076a0804-1527-4067-8d74-a1d0bb410da1_43.jpeg?w=768&q=80"

	u, err := url.Parse(urlString)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Wight:", u.Query()["w"])
	fmt.Println("Quality:", u.Query()["q"])
}
