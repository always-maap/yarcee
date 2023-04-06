package routes

import (
	"api/controllers"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	app.Get("/health_check", controllers.HealthController)

	AuthSetup(api)
}
