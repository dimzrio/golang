package main

import (
	"database/sql"
	"fmt"
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

func querySQL(db *sql.DB) ([]dimzrioRow, error) {
	row, err := db.Query("select * from dimzrio")
	defer row.Close()

	if err != nil {
		return nil, err
	}

	var result []dimzrioRow

	for row.Next() {
		var each = dimzrioRow{}
		err = row.Scan(&each.id, &each.name, &each.data)

		if err != nil {
			return nil, err
		}

		result = append(result, each)
	}

	return result, nil
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

	rowData, err := querySQL(db)

	if err != nil {
		log.Println(err)
	}

	for _, each := range rowData {
		fmt.Printf("ID: %d\tName: %s\tData: %s\n", each.id, each.name, each.data)
	}
}
