package controller

import (
	"api/broker"
	"api/database"
	"api/helper"
	"api/model"
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
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

type executeSandboxBody struct {
	Code string `json:"code"`
}

type sandboxJob struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

// @Summary  Update sandbox
// @Tags     Sandbox
// @Accept   json
// @Produce  json
// @Param    id      path int               true "id"
// @Param    request body executeSandboxBody true "query params"
// @Security Bearer
// @Router   /api/sandbox/{id}/execute [post]
func ExecuteSandbox(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(executeSandboxBody)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var existingSandbox model.Sandbox
	if err := database.DB.First(&existingSandbox, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"message": "Sandbox not found",
		})
	}

	if data.Code != existingSandbox.Code {
		existingSandbox.Code = data.Code

		if err := database.DB.Save(&existingSandbox).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"message": "Failed to update sandbox",
			})
		}
	}

	sandboxJob := &sandboxJob{
		ID:       existingSandbox.Id,
		Code:     existingSandbox.Code,
		Language: existingSandbox.Language,
	}

	sandboxJobJson, err := json.Marshal(sandboxJob)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to marshal message",
		})
	}

	publishSandboxJob("sandbox_job_ex", "sandbox_job_rk", string(sandboxJobJson))

	return c.JSON(&fiber.Map{
		"message": "success",
		"data":    existingSandbox,
	})
}

func publishSandboxJob(exchange, routingKey, body string) error {
	ch := broker.GetChannel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	return err
}
