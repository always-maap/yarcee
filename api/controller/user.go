package controller

import (
	"api/helper"

	"github.com/gofiber/fiber/v2"
)

// @Summary  User auth details
// @Tags     User
// @Accept   json
// @Produce  json
// @Security Bearer
// @Router   /api/user/ [get]
func RetrieveUserController(c *fiber.Ctx) error {
	user, err := helper.RetrieveUser(c.UserContext())

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
