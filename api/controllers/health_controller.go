package controllers

import "github.com/gofiber/fiber/v2"

func HealthController(c *fiber.Ctx) error {
	return c.SendString("OK")
}
