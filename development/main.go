package main

import "fmt"

// Family struct
type Family struct {
	Name    string
	Age     int
	Merried bool
}

// GetStatus func
func GetStatus(family []Family, status func(merried bool) string) []string {
	var result []string

	for _, data := range family {
		str := fmt.Sprintf("Name: %s\nAge: %d\nMerried: %s\n", data.Name, data.Age, status(data.Merried))
		result = append(result, str)
	}

	return result
}

func main() {
	data1 := Family{
		Name: "Dimas", Age: 28, Merried: true,
	}
	data2 := Family{
		Name: "Hana", Age: 28, Merried: true,
	}
	data3 := Family{
		Name: "Asiyah", Age: 1, Merried: false,
	}

	family := []Family{data1, data2, data3}

	result := GetStatus(family, func(merried bool) string {
		var status string
		if merried == true {
			status = "Merried"
		} else {
			status = "Single"
		}
		return status
	})

	// Print data
	for index, data := range result {
		fmt.Println("No. ", index)
		fmt.Println("Data: ", data)
	}

}
