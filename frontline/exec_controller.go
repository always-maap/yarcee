package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type ExecReq struct {
	ID       uint   `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

func execController(ctx *fiber.Ctx) error {
	execReq := new(ExecReq)

	if err := ctx.BodyParser(execReq); err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("/tmp/%d", execReq.ID))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	defer f.Close()

	_, err = f.WriteString(execReq.Code)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	switch execReq.Language {
	case "py":
		return py(ctx, execReq)
	case "node":
		return node(ctx, execReq)
	case "c":
		return cHandler(ctx, execReq)
	case "cpp":
		return cppHandler(ctx, execReq)
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Language not supported",
		})
	}
}
