package main

import (
	"github.com/gofiber/fiber/v2"                   // Import Fiber v2 correctly
	"github.com/gofiber/fiber/v2/middleware/cors"   // Correctly import CORS middleware
	"github.com/gofiber/fiber/v2/middleware/logger" // Import logger middleware from Fiber v2
	"main.go/database"
	"main.go/routers"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// Use the logger and CORS middleware
	app.Use(logger.New()) // Logging middleware
	app.Use(cors.New())   // CORS middleware

	// Setup routes
	routers.SetupRoutes(app)

	// Handle 404 - Not Found
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start the server on port 8080
	if err := app.Listen(":8080"); err != nil {
		panic(err) // Handle any error that occurs when starting the server
	}
}
