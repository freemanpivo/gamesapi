package server

import (
	"fmt"

	"github.com/freemanpivo/games-api/internal/config"
	"github.com/freemanpivo/games-api/internal/games"
	"github.com/freemanpivo/games-api/internal/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	// isolar logica dos handlers
	health.RegisterRoutes(app)
	games.RegisterRoutes(app)

	cfg := config.Get()
	app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))

	return app
}
