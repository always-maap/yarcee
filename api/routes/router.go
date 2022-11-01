package routes

import (
	"api/controllers"
	"api/database"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	database.Connect()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/health_check", controllers.HealthController)

	app.Post("/sign-up", controllers.SignUpController)
	app.Post("/sign-in", controllers.SignInController)
	app.Get("/sign-out", controllers.SignOutController)
	app.Get("/user", controllers.RetrieveUserController)
}
