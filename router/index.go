package router

import (
	mw "hgndgn/api/jwt-authentication/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

func Setup(app *fiber.App) {

	// use default Logger
	app.Use(middleware.Logger())

	// sets various HTTP headers for security
	app.Use(helmet.New())

	// public
	app.Group("/auth").
		Post("/login", Login)

	// middleware
	app.Use(mw.JwtMiddleware())

	// protected
	app.Get("/protected", FetchProtected)
}
