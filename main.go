package main

import (
	"github.com/chamanbravo/tinyurl/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")
	app.Post("/shorten", controllers.ShortenUrl)
	app.Get("/:url", controllers.Resolve)

	app.Listen(":3000")
}
