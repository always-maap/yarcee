package controller

import (
	"api/database"
	"api/helper"
	"api/model"

	"github.com/gofiber/fiber/v2"
)

func GetUserSandboxes(c *fiber.Ctx) error {
	user, err := helper.RetrieveUser(c.UserContext())

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var sandboxes []model.Sandbox
	database.DB.Where("user_refer = ?", user.Id).Find(&sandboxes)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    sandboxes,
	})

}

type createSandboxBody struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

// @Summary      Create sandbox
// @Tags         sandbox
// @Accept       json
// @Produce      json
// @Param request body createSandboxBody true "query params"
// @Router       /api/sandbox/ [post]
func CreateSandBox(c *fiber.Ctx) error {
	var data = new(createSandboxBody)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := helper.RetrieveUser(c.UserContext())

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sandbox := model.Sandbox{
		Name:      data.Name,
		Language:  data.Language,
		Code:      data.Code,
		UserRefer: user.Id,
	}

	database.DB.Create(&sandbox)

	return c.JSON(sandbox)
}
