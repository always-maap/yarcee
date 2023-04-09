package routes

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthSetup(api fiber.Router) {
	api.Post("/sign-up", controllers.SignUpController)
	api.Post("/sign-in", controllers.SignInController)
	api.Get("/user", controllers.RetrieveUserController)
}
