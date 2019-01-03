package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Data struct
type Data struct {
	ID   int    `json:"id"` // formating json to lowercase
	Name string `json:"name"`
}

func main() {
	data1 := Data{
		ID:   1,
		Name: "Dimas",
	}

	data2 := Data{
		ID:   2,
		Name: "Rio",
	}

	// Load to json
	d1, err := json.Marshal(data1)
	d2, err := json.Marshal(data2)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(d1) // byte string
	fmt.Println(d2) // byte string

	fmt.Println(string(d1))
	fmt.Println(string(d2))

	// Unload from json to struct

	var decode Data
	err = json.Unmarshal(d2, &decode)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(decode) // result {2 Rio}
	fmt.Println("ID: ", decode.ID)
	fmt.Println("Name: ", decode.Name)
}
