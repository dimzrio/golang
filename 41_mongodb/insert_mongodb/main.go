package main

import (
	"log"

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

	err = collection.Insert(&Person{Name: "dimas", Email: "dimas@gmail.com"},
		&Person{Name: "rio", Email: "rio@gmail.com"},
		&Person{Name: "tantowi", Email: "tantowi@gmail.com"})

	if err == nil {
		log.Println("Insert successfully...")
	}
}
