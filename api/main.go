package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Satvik-Srivastava/goshort/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1/shorten", routes.ShortenURL) // Better route
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found or error loading it")
	}

	// Get port from env with fallback
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // default port
	}
	if strings.HasPrefix(port, ":") {
		port = strings.TrimPrefix(port, ":")
	}

	app := fiber.New(fiber.Config{
		AppName: "GoShort - URL Shortener",
	})

	// Use Fiber's built-in logger middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}] ${status} - ${method} ${path}\n",
	}))

	setupRoutes(app)

	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
