package controllers

import (
	"api/database"
	"api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

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

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
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
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
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

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// @Summary      Sign out
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Router       /api/sign-out/ [get]
func SignOutController(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// @Summary      User auth details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Router       /api/user/ [get]
func RetrieveUserController(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("username = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}
