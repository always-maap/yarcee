package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SandboxSetup(api fiber.Router) {
	api.Post("/sandbox", middleware.Protected(), controllers.CreateSandBox)
	api.Get("/sandbox", middleware.Protected(), controllers.GetUserSandboxes)
}
