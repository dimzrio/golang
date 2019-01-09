package main

import "fmt"

// Family interface
type Family interface {
	getInfo() string
}

// Man struct
type Man struct {
	name, born string
	age        int
	merried    bool
}

// Lady struct
type Lady struct {
	name, born string
	age        int
	merried    bool
}

func (m *Man) getInfo() {
	var status string

	if m.merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}

	fmt.Printf(`*** MAN ***
name: %s
born: %s
age: %d
merried: %s`, m.name, m.born, m.age, status)
	fmt.Println()
}

func (l *Lady) getInfo() {
	var status string

	if l.merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}

	fmt.Printf(`*** LADY ***
name: %s
born: %s
age: %d
merried: %s`, l.name, l.born, l.age, status)
	fmt.Println()
}

func main() {
	person1 := &Man{name: "Rio", born: "Jakarta", age: 28, merried: true}
	person2 := &Lady{name: "Hana", born: "Cilacap", age: 28, merried: true}

	person1.getInfo()
	person2.getInfo()

}
