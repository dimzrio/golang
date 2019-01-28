package main

import (
	"math"
	"testing"
)

// Static result
var (
	// Initial Struct
	// cube           = Cube{Sisi: 4}
	// area   float64 = 96
	// around float64 = 48
	// volume float64 = 64

	cube           = Cube{Sisi: 2.3}
	area   float64 = 32
	around float64 = 28
	volume float64 = 12
)

func TestVolume(t *testing.T) {
	result := math.Round(cube.Volume())
	t.Log("Volume: ", result)
	if result != volume {
		t.Error("Error! No passed test, value", result)
	}
}

func TestAround(t *testing.T) {
	result := math.Round(cube.AroundCube())
	t.Log("Around: ", result)

	if result != around {
		t.Error("Error! No passed test, value", result)
	}
}

func TestArea(t *testing.T) {
	result := math.Round(cube.Area())
	t.Log("Area", result)

	if result != area {
		t.Error("Error! No passed test, value", result)
	}
}
