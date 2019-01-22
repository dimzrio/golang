package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// mysqladmin -u root password @xAmple123

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@xAmple123@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// format return (*sql.DB)
	// fmt.Printf("%T", db)

	return db, nil
}
func main() {
	db, err := connect()

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
}
