package main

import (
	"log"

	"github.com/danielmoisa/go-jwt/database"
	"github.com/danielmoisa/go-jwt/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Init db connection
	database.Connect()

	// Start fiber app
	app := fiber.New()

	// Set cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup routes
	routes.Setup(app)

	// Listen
	app.Listen(":4000")
}
