// package main

// import (
// 	"user/controllers"
// 	"user/models"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// 	"github.com/gofiber/fiber/v2/middleware/recover"
// )

// func main() {
// 	// Initialize Fiber app
// 	app := fiber.New()

// 	// Middleware
// 	app.Use(logger.New())
// 	app.Use(recover.New())

// 	// Connect to the database
// 	db, err := models.Database()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Initialize controllers
// 	userController := controllers.NewUserController(db)

// 	// Routes
// 	app.Get("/", controllers.Hello)
// 	app.Get("/err", controllers.Error)
// 	app.Get("/params:name", controllers.Params)
// 	app.Get("/header", controllers.Header)
// 	app.Get("/select", userController.Select)
// 	app.Post("/users", controllers.Post)
// 	app.Post("/insert", userController.Insert)
// 	app.Post("/delete", userController.Delete)
// 	app.Post("/update", userController.Update)

// 	// Start server
// 	app.Listen(":5000")
// }
