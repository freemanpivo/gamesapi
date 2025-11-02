package games

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers the /games route and wires repository -> service -> handler
// It expects the seed file to be at ./data/games_seed.json relative to the project root.
func RegisterRoutes(app *fiber.App) error {
	path := filepath.Join("data", "games_seed.json")
	repo, err := NewRepositoryFromFile(path)
	if err != nil {
		return err
	}
	service := NewService(repo)
	h := NewHandler(service)

	app.Get("/games", h.List)
	return nil
}
