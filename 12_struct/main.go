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

func (f Family) merrieds() bool {
	return f.merried
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
		age:     28,
		merried: true}

	fmt.Println(person1)
	fmt.Println(person1.names())
	fmt.Println(person1.merrieds())
	fmt.Println(person1.ages())

	// Anonymouse struct
	person2 := struct {
		name, gander string
		age          int
		merried      bool
	}{
		name:    "hana",
		gander:  "p",
		age:     28,
		merried: true}

	fmt.Println(person2)

	// assign value
	var person3 Family
	person3.name = "asiyah"
	person3.age = 1
	person3.merried = false
	person3.gander = "p"
	person3.brithday()
	person3.brithday()
	fmt.Println(person3)

}
