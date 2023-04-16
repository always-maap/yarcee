package controller

import "github.com/gofiber/fiber/v2"

// @Summary      Health check
// @Tags         Health check
// @Accept       json
// @Produce      json
// @Router       /health_check/ [get]
func HealthController(c *fiber.Ctx) error {
	return c.SendString("OK")
}
