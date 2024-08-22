package main

import (
	"GoFiber/config"
	"GoFiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the Fiber app
	app := fiber.New()

	// Initialize the database connection
	config.InitDatabase()

	// Set up routes
	routes.SetupRoutes(app)
	routes.SetupPostRoutes(app)

	// Start the server
	app.Listen(":3000")
}
