package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/developer", developer)
	app.Listen(":8080")
}

func developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
