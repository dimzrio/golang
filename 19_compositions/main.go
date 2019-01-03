package main

import "fmt"

/* compositions is instead inheritance oop in go */

// Biodata struct
type Biodata struct {
	FirstName, LastName, Gander string
	Age                         int
	Merried                     bool
}

// name fuct
func (b Biodata) name() string {
	return fmt.Sprintf("%s %s", b.FirstName, b.LastName)
}

// status func
func (b Biodata) status() string {
	var status string
	if b.Merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}
	return fmt.Sprintf(status)
}

// details func
func (b Biodata) details() string {
	status := b.status()
	return fmt.Sprintf(`*** Details ***
Name: %s %s
Age: %d
Status: %s`, b.FirstName, b.LastName, b.Age, status)
}

func main() {
	data := Biodata{"Dimas", "Riotantowi", "L", 28, true}
	fmt.Println(data.name())
	fmt.Println(data.details())
}
