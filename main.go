package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"

	"hgndgn/api/jwt-authentication/config"
	"hgndgn/api/jwt-authentication/jwt"
	"hgndgn/api/jwt-authentication/router"
)

func initialize() {
	const envFile = ".env"
	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}

	jwt.Initialize()
}

func main() {
	initialize()

	app := fiber.New()

	router.Setup(app)

	port := config.Get("port", "3000")

	app.Listen(port)
}
