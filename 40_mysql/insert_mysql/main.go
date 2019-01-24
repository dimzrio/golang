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

func insertSQL(db *sql.DB, id int, name, data string) error {
	_, err := db.Exec("insert into dimzrio value (?,?,?)", id, name, data)

	if err != nil {
		log.Fatalf("Failed to insert: %v\n", err)
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

	// Insert
	var id = 4
	var name = "hana"
	var data = "hello hana"

	err = insertSQL(db, id, name, data)

	if err != nil {
		log.Println(err)
	}

	log.Println("Insert successfully...")

}
