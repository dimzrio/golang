package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Details struct
type field struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
}

var data = []field{
	field{1, "Dimas", "Hello dimas"},
	field{2, "Rio", "Hello rio"},
}

func getUser(resp http.ResponseWriter, req *http.Request) {
	// Set reseponse apps/json
	resp.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if req.Method == "POST" {
		id, _ := strconv.Atoi(req.FormValue("id"))

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(resp, err.Error(), http.StatusInternalServerError)
					return
				}

				log.Println(req.RequestURI)
				resp.Write(result)
				return
			}
		}

		http.Error(resp, "Data not found", http.StatusBadRequest)
		log.Println(req.RequestURI, http.StatusBadRequest)
		return
	}
}

func getAllUser(resp http.ResponseWriter, req *http.Request) {
	// Set reseponse apps/json
	resp.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if req.Method == "POST" {
		result, err = json.Marshal(data)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(req.RequestURI)
		resp.Write(result)
		return
	}
}

func main() {
	// Index
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "This is an example json web api")
	})

	// User
	http.HandleFunc("/user", getUser)

	// Users
	http.HandleFunc("/users", getAllUser)

	// Lister web server
	http.ListenAndServe(":8080", nil)
}
