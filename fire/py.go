package main

import "github.com/gofiber/fiber/v2"

func py(c *fiber.Ctx) error {
	return c.SendString("py")
}
