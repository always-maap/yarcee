package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/health_check", healthCheckController)
	app.Post("/exec", execController)

	app.Listen(":8080")
}
