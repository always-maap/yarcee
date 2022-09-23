package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		createAndStartVmm(ctx)

		log.SetReportCaller(true)

		return c.SendString("yo")

	})

	app.Listen(":8081")
}
