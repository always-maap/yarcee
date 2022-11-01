package routes

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthSetup(v1 fiber.Router) {

	v1.Post("/sign-up", controllers.SignUpController)
	v1.Post("/sign-in", controllers.SignInController)
	v1.Get("/sign-out", controllers.SignOutController)
	v1.Get("/user", controllers.RetrieveUserController)
}
