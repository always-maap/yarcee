package routes

import (
	"api/controllers"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	app.Get("/health_check", controllers.HealthController)

	QuestionSetup(v1)

	AuthSetup(v1)
}
