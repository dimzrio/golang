package family

import (
	"fmt"
)

// Family struct
type Family struct {
	// First character of variable must Capitabize
	Name, Gander string
	Age          int
	Merried      bool
}

// New func
func New(name, gander string, age int, merried bool) Family {
	f := Family{Name: name, Gander: gander, Age: age, Merried: merried}
	return f
}

// Info func
func (f Family) Info() string {
	var status string
	if f.Merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}

	return fmt.Sprintf(`*** Info ***
Name: %s
Gander: %s
Age : %d
Merried : %s`, f.Name, f.Gander, f.Age, status)
}
