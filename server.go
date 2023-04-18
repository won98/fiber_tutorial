package main

import (
	"database/sql"
	"fmt"

	"gotest/controller"
	"gotest/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:1111@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Initialize User model and controller
	userModel := models.NewUserModel(db)
	userController := &controller.UserController{
		UserModel: userModel,
	}

	// Initialize Fiber app
	app := fiber.New()
	// Define routes
	app.Get("/", userController.Hello)
	app.Get("/error", userController.Error)
	app.Get("/params/:name", userController.Params)
	app.Get("/header", userController.Header)
	app.Post("/user", userController.Post)
	app.Post("/user/insert", userController.Insert)
	app.Post("/user/select", userController.Select)
	app.Post("/user/delete", userController.Delete)
	app.Post("/user/update", userController.Update)
	app.Post("/user/login", userController.Login)

	// Start server
	err = app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
