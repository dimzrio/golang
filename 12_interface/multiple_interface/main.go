package main

import "fmt"

// Blue interface
type Blue interface {
	carBlue() string
}

// Red interface
type Red interface {
	carRed() string
}

// Car struct
type Car struct {
	brand, year, price string
}

func (c Car) carBlue() string {
	return fmt.Sprintf(`*** Blue ***
Brand: %s
Year: %s
Price: %s`, c.brand, c.year, c.price)
}

func (c Car) carRed() string {
	return fmt.Sprintf(`*** Red ***
Brand: %s
Year: %s
Price: %s`, c.brand, c.year, c.price)
}

func main() {
	car := Car{brand: "BMW", year: "2006", price: "$50000"}

	var car1 Blue = car
	fmt.Println(car1.carBlue())

	var car2 Red = car
	fmt.Println(car2.carRed())

}
