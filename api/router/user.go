package router

import (
	"api/controller"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserSetup(api fiber.Router) {
	api.Get("/user", middleware.Protected(), controller.RetrieveUserController)
}
