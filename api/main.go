package main

import (
	"api/broker"
	"api/controller"
	"api/database"
	_ "api/docs"
	"api/router"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// @title                      YARCEE API
// @version                    1.0
// @description                YARCEE API swagger documentation
// @host                       localhost:8082
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.
//
// @BasePath                   /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Use(cors.New())

	router.Setup(app)

	if err := broker.Connect(); err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range broker.GetSandboxStatusJobsChan() {
			var jobStatus controller.JobStatus
			fmt.Println(string(d.Body))
			err := json.Unmarshal([]byte(d.Body), &jobStatus)
			if err != nil {
				fmt.Println("Received invalid job")
				continue
			}
			controller.SandboxJobStatusHandler(jobStatus)
		}
	}()

	app.Listen(":8082")
}
