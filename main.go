package main

import (
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Name  string `form:"name"`
	Image []byte `form:"image"`
}

func main() {
	app := fiber.New()

	app.Static("/", "index.html")
	app.Listen(":3000")
}
