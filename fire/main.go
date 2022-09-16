package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/health_check", healthCheck)
	app.Get("/run", py)

	app.Listen(":8080")
}

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
