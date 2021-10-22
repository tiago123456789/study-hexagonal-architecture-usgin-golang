package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	dbModule "github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/adapters/db"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

const (
	host     = "database"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "test"
)

func createTableUser(db *sql.DB) {
	sql := "CREATE TABLE IF NOT EXISTS products("
	sql += " id VARCHAR(255), "
	sql += " name VARCHAR(150), "
	sql += " price DECIMAL(10, 2), "
	sql += " status VARCHAR(50) "
	sql += " );"
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func setUp() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	createTableUser(db)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := setUp()
	productDb := dbModule.NewProductDb(db)
	productService := application.NewProductService(productDb)
	productCreated, err := productService.Create("product abc", 10.0)
	if err != nil {
		fmt.Println(err)
	}

	product, err := productService.Get(productCreated.GetID())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v", product)
}
