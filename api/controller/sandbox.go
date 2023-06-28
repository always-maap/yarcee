package controller

import (
	"api/database"
	"api/helper"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

// @Summary  Get user sandboxes
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Security Bearer
// @Router   /api/sandbox/ [get]
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

// @Summary  Get sandbox
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Param    id path int true "id"
// @Security Bearer
// @Router   /api/sandbox/{id} [get]
func GetSandbox(c *fiber.Ctx) error {
	id := c.Params("id")

	var sandbox model.Sandbox
	database.DB.First(&sandbox, id)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    sandbox,
	})
}

type createSandboxBody struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

// @Summary  Create sandbox
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Param    request body createSandboxBody true "query params"
// @Security Bearer
// @Router   /api/sandbox/ [post]
func CreateSandbox(c *fiber.Ctx) error {
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

type updateSandboxBody struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

// @Summary  Update sandbox
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Param    id      path int               true "id"
// @Param    request body updateSandboxBody true "query params"
// @Security Bearer
// @Router   /api/sandbox/{id} [put]
func UpdateSandbox(c *fiber.Ctx) error {
	id := c.Params("id")
	var data = new(updateSandboxBody)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var sandbox model.Sandbox
	database.DB.Model(&sandbox).Clauses(clause.Returning{}).Where("id = ?", id).Updates(model.Sandbox{Name: data.Name, Language: data.Language, Code: data.Code})

	return c.JSON(&fiber.Map{
		"message": "success",
		"data":    sandbox,
	})
}

// @Summary  Delete sandbox
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Param    id path int true "id"
// @Security Bearer
// @Router   /api/sandbox/ [delete]
func DeleteSandbox(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&model.Sandbox{}, id)

	return c.JSON(&fiber.Map{
		"message": "success",
	})
}
