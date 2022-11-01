package routes

import (
	"api/controllers"

	"github.com/gofiber/fiber/v2"
)

func QuestionSetup(v1 fiber.Router) {
	v1.Get("/questions", controllers.GetAllQuestion)
	v1.Get("/questions/:id", controllers.GetQuestion)
}
