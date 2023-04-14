package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthSetup(api fiber.Router) {
	api.Post("/sign-up", controllers.SignUpController)
	api.Post("/sign-in", controllers.SignInController)
	api.Get("/user", middleware.Protected(), controllers.RetrieveUserController)
}
