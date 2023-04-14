package middleware

import (
	"api/database"
	"api/models"
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
)

func Protected() fiber.Handler {
	secret := os.Getenv("SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(secret),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func jwtSuccess(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	jwtToken := strings.Split(headers["Authorization"], " ")[1]

	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("username = ?", claims.Issuer).First(&user)

	c.SetUserContext(context.WithValue(c.UserContext(), "user", user))

	return c.Next()
}
