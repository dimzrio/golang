package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Formater struct
type Formater struct {
	ID   int
	Name string
	Data string
}

func getData() ([]Formater, error) {
	var err error
	var client = &http.Client{}
	var data []Formater

	req, err := http.NewRequest("POST", "http://localhost:8080/users", nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func main() {
	var user, _ = getData()

	for _, each := range user {
		out := fmt.Sprintf("ID: %d\tName: %s\t Data: %s", each.ID, each.Name, each.Data)
		fmt.Println(out)
	}
}
