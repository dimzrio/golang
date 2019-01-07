package main

import (
	"fmt"
)

func main() {
	cat := []string{"Kuro", "Ema", "Nyiung", "Centimeters"}
	cat = append(cat, "Emy")               // Append data
	fmt.Println("Length:", len(cat))       // Length array
	fmt.Println("Capacity cat:", cap(cat)) // Capacity array, default 8
	fmt.Println("Data:", cat[:])           // Get all data
	fmt.Println("Data [0-2]:", cat[0:2])   // Kuro, Ema --> start index 0, stop before index 2
	fmt.Println("Data [2:]:", cat[2:])     // Get all data after index 2
	fmt.Println("Data [:-1]:", cat[:4])    // Get all data before index 3
}
