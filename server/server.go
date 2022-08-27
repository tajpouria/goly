package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	RegisterGolyRoute(app.Group("/v1/goly"))
	RegisterRedirectRouter(app.Group("/r"))

	app.Listen(":8080")
}
