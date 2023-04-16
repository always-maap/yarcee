package router

import (
	"api/controller"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SandboxSetup(api fiber.Router) {
	api.Post("/sandbox", middleware.Protected(), controller.CreateSandBox)
	api.Get("/sandbox", middleware.Protected(), controller.GetUserSandboxes)
}
