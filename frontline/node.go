package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func node(c *fiber.Ctx, execReq *ExecReq) error {
	return execCmd(c, "/usr/bin/node", fmt.Sprintf("/tmp/%d", execReq.ID))
}
