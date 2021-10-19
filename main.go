package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTableUser(db *sql.DB) {
	sql := "CREATE TABLE IF NOT EXISTS user("
	sql += " username VARCHAR(255), "
	sql += " password VARCHAR(255) );"
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	createTableUser(db)
	stmt, err := db.Prepare("INSERT INTO user(username, password) values(?,?)")
	_, err = stmt.Exec("tiago", "tiago123456789")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted success")
}
