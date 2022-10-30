package main

import (
	"github.com/gofiber/swagger"

	_ "github.com/always-maap/yarcee/api/docs"
	"github.com/gofiber/fiber/v2"
)

// @title          Fiber Example API
// @version        1.0
// @description    This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  fiber@swagger.io
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:3000
// @BasePath       /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
