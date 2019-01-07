package main

import "fmt"

func main() {

	// Example 1
	color := "red"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue", "purple":
		fmt.Println("Color is blue or purple")
	default:
		fmt.Println("Color is not blue or red")
	}

	// Example 2
	point := 0
	switch {
	case point >= 7:
		fmt.Println("Great")
	case (point >= 5) && (point <= 6):
		fmt.Println("Enough")
		fallthrough // This func force to run next case
	case (point > 2) && (point < 4):
		fmt.Println("Not bad, You need learn more")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn more and remedial")
		}

	}
}
