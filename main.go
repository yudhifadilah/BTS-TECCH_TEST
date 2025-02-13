package main

import (
	"log"
	"techtest/configs"
	"techtest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Connect to the Database
	configs.ConnectDatabase()

	// Check database connection
	if configs.DB == nil {
		log.Fatal("Failed to connect to the database")
	} else {
		log.Println("Database connected successfully")
	}

	// Start a new fiber app
	app := fiber.New()

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Ganti dengan domain yang diperbolehkan
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	// Middleware Logging dan Recovery
	app.Use(logger.New())
	app.Use(recover.New())

	// Setup Routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
