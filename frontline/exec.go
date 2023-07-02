package main

import (
	"bytes"
	"net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func execCmd(c *fiber.Ctx, prog string, args ...string) error {
	var execStdOut, execStdErr bytes.Buffer

	cmd := exec.Command(prog, args...)
	cmd.Stdout = &execStdOut
	cmd.Stderr = &execStdErr

	err := cmd.Run()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"stdout": execStdOut.String(),
			"stderr": execStdErr.String(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "OK",
		"stdout":  execStdOut.String(),
		"stderr":  execStdErr.String(),
	})
}
