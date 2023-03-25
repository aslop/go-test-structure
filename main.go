package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type AppInstance struct {
	App *fiber.App
}

func init() {
}

func main() {
	app := fiber.New()

	// ==================================================================
	// CORS
	// ==================================================================
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// ==================================================================
	// Middlewares
	// ==================================================================
	app.Use(logger.New())
	app.Use(requestid.New())

	// ==================================================================
	// Routes
	// ==================================================================
	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5005")

}
