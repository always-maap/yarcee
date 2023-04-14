package controllers

import (
	"api/database"
	"api/helper"
	"api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func checkPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

type signUpBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary      Sign up
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body signUpBody true "query params"
// @Router       /api/sign-up/ [post]
func SignUpController(c *fiber.Ctx) error {
	var data = new(signUpBody)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := hashPassword(data.Password)
	user := models.User{
		Name:     data.Name,
		Username: data.Username,
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)

}

type signInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary      Sign in
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body signInBody true "query params"
// @Router       /api/sign-in/ [post]
func SignInController(c *fiber.Ctx) error {
	var data = new(signInBody)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("username = ?", data.Username).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if !checkPasswordHash(data.Password, user.Password) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Username,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    token,
	})
}

// @Summary      User auth details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security Bearer
// @Router       /api/user/ [get]
func RetrieveUserController(c *fiber.Ctx) error {
	user, err := helper.RetrieveUser(c.UserContext())

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
