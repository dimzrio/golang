package main

import "fmt"

// init fucn run before than main func running

// Private struct
type info struct {
	name, email string
}

var inf info

func (i info) GetInfo() {
	fmt.Printf("*** INFO ***\nName: %s\nEmail: %s\n", i.name, i.email)
}

func init() {
	inf.name = "Mo. Salah"
	inf.email = "salah@liverpool.en"
	fmt.Println("This is an func init example")
}

func main() {
	inf.GetInfo()
}
