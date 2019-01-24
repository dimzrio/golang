package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dimzrioRow struct {
	id   int
	name string
	data string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@xAmple123@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func insertSQL(db *sql.DB, name string) error {
	_, err := db.Exec("update dimzrio set data=? where name=?", "hello world", name)

	if err != nil {
		log.Fatalf("Failed to update: %v\n", err)
		return err
	}

	return nil
}

func main() {
	db, err := connect()

	if err != nil {
		log.Println(err)
		return
	}
	err = db.Ping()

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connecting to db successfully..")

	// Update
	var name = "hana"

	err = insertSQL(db, name)

	if err != nil {
		log.Println(err)
	}

	log.Println("Update successfully...")

}
