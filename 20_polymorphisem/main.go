package main

import (
	"fmt"
	"math"
)

// Volume interface
type Volume interface {
	vol() float64 // type must be define
}

// Cubes struct
type Cubes struct {
	// Rumus sisi x sisi x sisi
	sisi float64
}

// Tubes struct
type Tubes struct {
	// Rumus phi x radius x radius x tinggi
	radius, tinggi float64
}

func (c Cubes) vol() float64 {
	volume := c.sisi * c.sisi * c.sisi
	return volume
}

func (t Tubes) vol() float64 {
	volume := math.Phi * (t.radius * t.radius * t.tinggi)
	return volume
}

func getVolume(v []Volume) {
	for _, getResult := range v {
		fmt.Printf("Volumes: %.2f\n", getResult.vol())
	}
}

func main() {
	calc1 := Cubes{sisi: 5}
	calc2 := Tubes{radius: 6, tinggi: 7}
	result := []Volume{calc1, calc2}
	getVolume(result)
}
