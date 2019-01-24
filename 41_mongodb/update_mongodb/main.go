package main

import (
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

	var selector = bson.M{"name": "tantowi"}
	var change = Person{"tantowi", "tantowi@yahoo.com"}

	err = collection.Update(selector, change)

	if err == nil {
		log.Println("Update successfully...")
	}
}
