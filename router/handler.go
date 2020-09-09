package router

import (
	JWT "hgndgn/api/jwt-authentication/jwt"
	. "hgndgn/api/jwt-authentication/model"

	jwt "github.com/dgrijalva/jwt-go"

	"log"

	"github.com/gofiber/fiber"
)

func Login(c *fiber.Ctx) {
	var res JSONResponse
	var loginData LoginCredentials
	c.BodyParser(&loginData)

	if err := c.BodyParser(&loginData); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(res.WithMessage("Invalid body").WithData(err))
		return
	}

	if loginData.Username == "admin" && loginData.Password == "123" {
		token := JWT.CreateToken(loginData.Username)
		tokenString := JWT.TokenString(token)
		expiresAt := JWT.TokenClaims(token).ExpiresAt

		res = res.WithData(LoginResponse{Token: tokenString, ExpiresAt: expiresAt}).Build()
		c.JSON(res.Data)
		log.Println(res.Data)
		return
	}

	res = res.WithMessage("InvalidCredentials").Build()
	c.Status(fiber.StatusBadRequest).JSON(res)
}

func FetchProtected(c *fiber.Ctx) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	c.Send("Hi " + username)
}
