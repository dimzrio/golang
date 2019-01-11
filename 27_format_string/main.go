package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	fmt.Printf("1. %b\n", data.age)             //Numeric biner
	fmt.Printf("2. %o\n", data.age)             // Numeric base 8 octa
	fmt.Printf("3. %d\n", data.age)             // Numeric base 10 decimal
	fmt.Printf("4. %x\n", data.age)             // Numeric base 16 hexa
	fmt.Printf("5. %c\n", 1400)                 // Unicode character
	fmt.Printf("6. %e\n", data.height)          // Scientific notation
	fmt.Printf("7. %e\n", data.height)          //Scientific notation
	fmt.Printf("8. %f\n", data.height)          // Float 6 digit
	fmt.Printf("9. %.f\n", data.height)         // Float without digit
	fmt.Printf("10. %.2f\n", data.height)       // Float 2 digit
	fmt.Printf("11. %p\n", &data.name)          // Pointer with prefix 0x
	fmt.Printf("12. %q\n", `" name \ height "`) // Escape string
	fmt.Printf("13. %s\n", data.name)           // String
	fmt.Printf("14. %t\n", data.isGraduated)    // Bool
	fmt.Printf("15. %T\n", data.height)         // Type of variable
	fmt.Printf("16. %v\n", data)                // Acceptable all format, include type interface{} or time
	fmt.Printf("17. %+v\n", data)               // Struct
	fmt.Printf("18. %#v\n", data)               // Struct with field
	fmt.Printf("19. %%\n")                      // Character %

}
