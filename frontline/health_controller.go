package main

import "github.com/gofiber/fiber/v2"

func healthCheckController(c *fiber.Ctx) error {
	return c.SendString("OK")
}
