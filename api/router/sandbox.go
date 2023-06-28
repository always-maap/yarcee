package router

import (
	"api/controller"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SandboxSetup(api fiber.Router) {
	api.Get("/sandbox", middleware.Protected(), controller.GetUserSandboxes)
	api.Get("/sandbox/:id", middleware.Protected(), controller.GetSandbox)
	api.Post("/sandbox", middleware.Protected(), controller.CreateSandbox)
	api.Put("/sandbox/:id", middleware.Protected(), controller.UpdateSandbox)
	api.Delete("/sandbox/:id", middleware.Protected(), controller.DeleteSandbox)
}
