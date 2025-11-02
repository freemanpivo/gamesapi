package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/freemanpivo/games-api/internal/games"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// mount feature routes
	games.RegisterRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
