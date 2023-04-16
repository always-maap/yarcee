package router

import (
	"api/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthSetup(api fiber.Router) {
	api.Post("/sign-up", controller.SignUpController)
	api.Post("/sign-in", controller.SignInController)
}
