package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// Person struct
type Person struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func main() {
	var conn, err = mgo.Dial("mongodb://rio:detikc0m@localhost")
	defer conn.Close()

	if err != nil {
		log.Fatalf("Failed connect to db, Error: %v", err)
	}

	collection := conn.DB("golang").C("example")

	if err != nil {
		log.Fatalf("Failed connect to collection, Error: %v", err)
	}

	var result = Person{}
	var selector = bson.M{"name": "dimas"}
	err = collection.Find(selector).One(&result)

	fmt.Println("Name:", result.Name)
	fmt.Println("Email:", result.Email)

	// Sparate
	fmt.Println()

	// get all data
	var data []Person

	err = collection.Find(nil).All(&data)

	for _, each := range data {
		fmt.Println(each)
	}
}
