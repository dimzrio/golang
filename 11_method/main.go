package main

import "fmt"

// Family structure
type Family struct {
	name, gander string
	age          int
	merried      bool
}

func (f Family) names() string {
	return f.name
}

func (f Family) merrieds() string {
	var status string

	if f.merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}

	return status
}

func (f Family) ages() int {
	return f.age
}

func (f *Family) brithday() {
	f.age++
}

func main() {

	person1 := Family{
		name:    "Rio",
		gander:  "l",
		age:     27,
		merried: true,
	}

	fmt.Println("Name:", person1.names())
	fmt.Println("Status:", person1.merrieds())

	person1.brithday() // Add +1 for age

	fmt.Println("Age:", person1.ages())

}
