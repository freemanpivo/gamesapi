package games

import (
	"github.com/gofiber/fiber/v2"
)

type GamesResponse struct {
	Data []Game `json:"data"`
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

type Handler struct {
	service *Service
}

func (h *Handler) List(c *fiber.Ctx) error {
	name := c.Query("name", "")
	platform := c.Query("platform", "")
	genre := c.Query("genre", "")
	subGenre := c.Query("subGenre", "")

	games, err := h.service.GetGames(name, platform, genre, subGenre)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(GamesResponse{Data: games})
}
