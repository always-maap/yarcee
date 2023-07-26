package main

import (
	"bytes"
	"net/http"
	"os/exec"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ExecRes struct {
	Message      string `json:"message"`
	StdErr       string `json:"stderr"`
	StdOut       string `json:"stdout"`
	ExecDuration int64  `json:"exec_duration"`
	MemUsage     int64  `json:"mem_usage"`
}

func execCmd(c *fiber.Ctx, prog string, args ...string) error {
	var execStdOut, execStdErr bytes.Buffer

	cmd := exec.Command(prog, args...)
	cmd.Stdout = &execStdOut
	cmd.Stderr = &execStdErr

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ExecRes{
			Message:      "Failed to exec",
			StdOut:       execStdOut.String(),
			StdErr:       execStdErr.String(),
			ExecDuration: elapsed.Milliseconds(),
			MemUsage:     cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss,
		})
	}

	return c.Status(http.StatusOK).JSON(ExecRes{
		Message:      "Success",
		StdOut:       execStdOut.String(),
		StdErr:       execStdErr.String(),
		ExecDuration: elapsed.Milliseconds(),
		MemUsage:     cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss,
	})
}
