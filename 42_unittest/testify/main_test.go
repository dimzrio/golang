package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Static result
var (
	cube           = Cube{Sisi: 2.3}
	area   float64 = 32
	around float64 = 28
	volume float64 = 12
)

func TestVolume(t *testing.T) {
	assert.Equal(t, math.Round(cube.Volume()), volume, "No Passed")
}

func TestArea(t *testing.T) {
	assert.Equal(t, math.Round(cube.Area()), area, "No Passed")
}

func TestArround(t *testing.T) {
	assert.Equal(t, math.Round(cube.AroundCube()), around, "No Passed")
}
