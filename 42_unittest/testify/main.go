package main

// go test main.go main_test.go
// go get github.com/stretchr/testify

import (
	"fmt"
)

// Cube struct
type Cube struct {
	Sisi float64
}

// Volume func method
func (k Cube) Volume() float64 {
	return float64(k.Sisi * k.Sisi * k.Sisi)
}

// AroundCube func method
func (k Cube) AroundCube() float64 {
	return 12 * k.Sisi
}

// Area func method
func (k Cube) Area() float64 {
	return (k.Sisi * k.Sisi) * 6
}

func main() {

	cube := Cube{Sisi: 2.3}
	fmt.Println(cube.Volume())
	fmt.Println(cube.AroundCube())
	fmt.Println(cube.Area())
}
