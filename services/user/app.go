package main

import (
	"flag"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"user/database"
	"user/handlers"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Bind handlers
	app.Get("/users", handlers.UserList)
	app.Post("/users", handlers.UserCreate)
	app.Get("/healthz", handlers.Health)

	// Handle not founds
	app.Use(handlers.NotFound)

	// Put flags `-port=:8080` for example to start listening in port 8080
	log.Fatal(app.Listen(*port))
}
