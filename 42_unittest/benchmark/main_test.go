package main

import "testing"

var cube = Cube{Sisi: 4}

func BenchmarkVolume(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		cube.Volume()
	}
}
