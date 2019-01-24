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

func querySQL(db *sql.DB, id int) (dimzrioRow, error) {
	smt, err := db.Prepare("select * from dimzrio where id=?")

	if err != nil {
		log.Fatalf("Failed to select row: %v\n", err)
	}

	var result dimzrioRow
	err = smt.QueryRow(id).Scan(&result.id, &result.name, &result.data)

	if err != nil {
		log.Fatalf("Failed to get data: %v\n", err)
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

	id := 3
	rowData, err := querySQL(db, id)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("ID: %d\tName: %s\tData: %s\n", rowData.id, rowData.name, rowData.data)

	id = 2
	rowData, err = querySQL(db, id)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("ID: %d\tName: %s\tData: %s\n", rowData.id, rowData.name, rowData.data)

}
