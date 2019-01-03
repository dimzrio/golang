package main

import "fmt"

// PersonalInformation interface
type PersonalInformation interface {
	employeData() string
}

// SalaryEmployee interface
type SalaryEmployee interface {
	employeeSalary() string
}

// DetailsEmployee interface
type DetailsEmployee interface {
	PersonalInformation
	SalaryEmployee
}

// Employee struct
type Employee struct {
	name, born, gander string
	age                int
	merried            bool
	salary             float64
}

//PersonalInformation functions
func (employee Employee) employeData() string {
	var status string

	if employee.merried == true {
		status = "Merried"
	} else {
		status = "Single"
	}

	return fmt.Sprintf(`Name: %s
Born: %s
Gander: %s
Age: %d
Status: %s`, employee.name, employee.born, employee.gander, employee.age, status)
}

//SalaryEmployee functions
func (employee Employee) employeeSalary() string {
	return fmt.Sprintf("Salary: $%.2f", employee.salary)
}

func main() {
	employee := Employee{
		name:    "Dimas Rio",
		born:    "Jakarta",
		gander:  "Male",
		age:     28,
		merried: true,
		salary:  800.0}

	var details DetailsEmployee = employee
	fmt.Println(details.employeData())
	fmt.Println(details.employeeSalary())
}
