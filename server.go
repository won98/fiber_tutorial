package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/your-username/your-project/controllers"
)

func Database() (*sql.DB, error) {
	// ...
}

func main() {
	app := fiber.New()

	app.Get("/", controllers.Hello)
	app.Get("/err", controllers.Error)
	app.Get("/params:name", controllers.Params)
	app.Get("/header", controllers.Header)
	app.Get("/select", controllers.Select)
	app.Post("/users", controllers.Post)
	app.Post("/insert", controllers.Insert)
	app.Post("/delete", controllers.Delete)
	app.Post("/update", controllers.Update)

	app.Listen(":5000")
}
