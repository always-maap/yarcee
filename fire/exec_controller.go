package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type ExecReq struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

func execController(c *fiber.Ctx) error {
	execReq := new(ExecReq)

	if err := c.BodyParser(execReq); err != nil {
		return err
	}

	f, err := os.Create("/tmp/" + execReq.Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	defer f.Close()

	_, err = f.WriteString(execReq.Code)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	switch execReq.Language {
	case "py":
		return py(c, execReq)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Language not supported",
		})
	}
}
