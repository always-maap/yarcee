package controllers

import (
	"api/database"
	"api/models"

	"github.com/gofiber/fiber/v2"
)

type QuestionList struct {
	Id         uint   `json:"id"`
	No         int    `json:"no"`
	Name       string `json:"name"`
	Subject    string `json:"subject"`
	Difficulty string `json:"difficulty"`
}

func GetAllQuestion(c *fiber.Ctx) error {
	var questions []QuestionList
	database.DB.Model(&models.Question{}).Find(&questions)
	return c.JSON(questions)
}

func GetQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	var question models.Question
	database.DB.Model(&models.Question{}).Where("id = ?", id).First(&question)
	return c.JSON(question)
}
