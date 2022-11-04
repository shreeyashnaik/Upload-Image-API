package main

import (
	"log"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Name  string                `form:"name"`
	Image *multipart.FileHeader `form:"image"`
}

func main() {
	app := fiber.New()

	app.Post("/upload_image", func(ctx *fiber.Ctx) error {
		p := new(Payload)
		if err := ctx.BodyParser(p); err != nil {
			log.Println("image upload error --> ", err)
			return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}

		ctx.SaveFile(p.Image, "./images/1.png")

		return ctx.JSON(fiber.Map{"status": 200, "msg": "Successfully uploaded"})
	})

	app.Listen(":3000")
}
