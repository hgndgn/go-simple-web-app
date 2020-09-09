package middleware

import (
	"hgndgn/api/jwt-authentication/jwt"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func JwtMiddleware() func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwt.JwtConfig.Secret),
	})
}
