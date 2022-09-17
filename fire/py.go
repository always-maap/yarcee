package main

import "github.com/gofiber/fiber/v2"

func py(c *fiber.Ctx, execReq *ExecReq) error {
	return execCmd(c, "python3", execReq.Id)
}
