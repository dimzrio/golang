package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dimzrioTable struct {
	Name string
	Data string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@xAmple123@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func selectDB(db *sql.DB, id int) error {
	log.Println("Selecting table dimzrio...")
	var result = dimzrioTable{}

	err := db.QueryRow("select name, data from dimzrio where id=?", id).Scan(&result.Name, &result.Data)

	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", result)

	return nil
}

func main() {
	db, err := connect()
	defer db.Close()

	if err != nil {
		log.Println(err)
		return
	}
	err = db.Ping()

	if err != nil {
		log.Panicln(err)
		return
	}
	log.Println("Connecting to db successfully..")

	// Select
	id := 2
	err = selectDB(db, id)

}
