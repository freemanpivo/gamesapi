package games

import (
	"github.com/gofiber/fiber/v2"
)

// response wrapper
type GamesResponse struct {
	Data []Game `json:"data"`
}

// NewHandler constructs a Fiber handler with service injected
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

type Handler struct {
	service *Service
}

func (h *Handler) List(c *fiber.Ctx) error {
	// query params: name, platform, gender, subGender
	name := c.Query("name", "")
	platform := c.Query("platform", "")
	gender := c.Query("gender", "")
	subGender := c.Query("subGender", "")

	games, err := h.service.GetGames(name, platform, gender, subGender)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(GamesResponse{Data: games})
}
