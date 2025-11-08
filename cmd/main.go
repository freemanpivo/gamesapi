package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/freemanpivo/games-api/internal/games"
	"github.com/freemanpivo/games-api/internal/health"
)

func NewApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	health.RegisterRoutes(app)
	games.RegisterRoutes(app)

	return app
}

func main() {
	app := NewApp()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
