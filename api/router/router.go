package router

import (
	"api/controller"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	app.Get("/health_check", controller.HealthController)

	AuthSetup(api)
	SandboxSetup(api)
	UserSetup(api)
}
