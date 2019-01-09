package main

import "fmt"

// Animal interface
type Animal interface {
	info() string
}

// Cat struct
type Cat struct {
	name string
	age  int
}

// Dog struct
type Dog struct {
	name string
	age  int
}

func (c Cat) info() string {
	return fmt.Sprintf(`*** CAT ***
name: %s
age: %d
`, c.name, c.age)
}

func (d Dog) info() string {
	return fmt.Sprintf(`*** DOG ***
name: %s
age: %d
`, d.name, d.age)
}

func getInfo(a []Animal) {
	for _, animal := range a {
		fmt.Println(animal.info())
	}
}

func main() {
	nyiung := Cat{name: "Nyiung", age: 1}
	ema := Cat{name: "Ema", age: 4}
	kuro := Cat{name: "Kuro", age: 4}
	shiro := Dog{name: "Shiro", age: 5}
	animal := []Animal{nyiung, ema, kuro, shiro}
	getInfo(animal)
}
