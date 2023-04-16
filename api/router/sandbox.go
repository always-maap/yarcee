package router

import (
	"api/controller"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SandboxSetup(api fiber.Router) {
	api.Post("/sandbox", middleware.Protected(), controller.CreateSandbox)
	api.Get("/sandbox", middleware.Protected(), controller.GetUserSandboxes)
	api.Put("/sandbox/:id", middleware.Protected(), controller.UpdateSandbox)
	api.Delete("/sandbox/:id", middleware.Protected(), controller.DeleteSandbox)
}
