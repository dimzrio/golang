package main

import "fmt"

func main() {

	data := map[string]interface{}{
		"name":    "Mo. Salah",
		"age":     28,
		"height":  175.5,
		"merried": true,
		"history": []string{"Chelsea", "As Roma", "Liverpool"},
	}

	// Manual
	// fmt.Println(data["name"].(string))
	// fmt.Println(data["age"].(int))
	// fmt.Println(data["height"].(float64))
	// fmt.Println(data["merried"].(bool))
	// fmt.Println(data["history"].([]string))

	// Auto
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
