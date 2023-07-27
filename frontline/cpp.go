package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func cppHandler(ctx *fiber.Ctx, execReq *ExecReq) error {
	var compileStdOut, compileStdErr bytes.Buffer
	compileCmd := exec.Command("g++", "-x", "c++", fmt.Sprintf("/tmp/%d", execReq.ID), "-o", fmt.Sprintf("/tmp/%d.out", execReq.ID))
	compileCmd.Stdout = &compileStdOut
	compileCmd.Stderr = &compileStdErr
	err := compileCmd.Run()

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ExecRes{
			Message: "Failed to compile",
			StdOut:  compileStdOut.String(),
			StdErr:  compileStdErr.String(),
		})
	}

	return execCmd(ctx, fmt.Sprintf("/tmp/%d.out", execReq.ID))
}
