package main

import "fmt"

func details(name, email string, age int) string {
	return fmt.Sprintf(`
*** Info ***
name : %s
age : %d
email : %s`, name, age, email)
}

func main() {
	fmt.Println(details("Dimas", "dimas@gmail.com", 28))
}
