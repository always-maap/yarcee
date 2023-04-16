package routes

import (
	"api/controller"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthSetup(api fiber.Router) {
	api.Post("/sign-up", controller.SignUpController)
	api.Post("/sign-in", controller.SignInController)
	api.Get("/user", middleware.Protected(), controller.RetrieveUserController)
}
