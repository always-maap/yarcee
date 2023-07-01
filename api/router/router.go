package router

import (
	"api/controller"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/health_check", controller.HealthController)

	api := app.Group("/api")

	AuthSetup(api)
	SandboxSetup(api)
	UserSetup(api)
}
