package web

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/adapters/db"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/adapters/web/handlers"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

type WebServer struct {
	app *fiber.App
}

func getConnection() *sql.DB {
	url := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"database", 5432, "postgres", "root", "test",
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("No possible connect in database")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sql := "CREATE TABLE IF NOT EXISTS products("
	sql += " id varchar(255), "
	sql += " name varchar(80), "
	sql += " price decimal(10, 2), "
	sql += " status varchar(80) ) "

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func (w *WebServer) LoadRoutes() {
	w.app = fiber.New()
	productRepository := db.NewProductDb(getConnection())
	productService := application.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)
	w.app.Get("/:id", productHandler.GetById)
	w.app.Post("/", productHandler.Save)

}

func (w *WebServer) Start() {
	w.app.Listen(":3000")
}
