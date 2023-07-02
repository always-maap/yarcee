package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func py(c *fiber.Ctx, execReq *ExecReq) error {
	return execCmd(c, "/usr/bin/python3", fmt.Sprintf("/tmp/%d", execReq.ID))
}
