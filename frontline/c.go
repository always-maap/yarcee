package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func cHandler(c *fiber.Ctx, execReq *ExecReq) error {
	var compileStdOut, compileStdErr bytes.Buffer
	compileCmd := exec.Command("gcc", "-x", "c", fmt.Sprintf("/tmp/%d", execReq.ID), "-o", fmt.Sprintf("/tmp/%d.out", execReq.ID))
	compileCmd.Stdout = &compileStdOut
	compileCmd.Stderr = &compileStdErr
	err := compileCmd.Run()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ExecRes{
			Message: "Failed to compile",
			StdOut:  compileStdOut.String(),
			StdErr:  compileStdErr.String(),
		})
	}

	return execCmd(c, fmt.Sprintf("/tmp/%d.out", execReq.ID))
}
