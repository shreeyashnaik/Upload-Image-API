package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Name  string `form:"name"`
	Image string `form:"image"`
}

func main() {
	app := fiber.New()

	app.Post("/upload_image", func(ctx *fiber.Ctx) error {
		file, _ := ctx.FormFile("image")

		ctx.SaveFile(file, "./images/1.png")

		p := new(Payload)
		ctx.BodyParser(p)

		log.Println("Name: ", p.Name, " Image: ", p.Image)
		return nil
	})

	app.Listen(":3000")
}
