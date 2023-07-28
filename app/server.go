package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./app/views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./app/static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Listen(":3030")
}
