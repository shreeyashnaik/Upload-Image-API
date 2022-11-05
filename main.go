package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/upload_image", func(ctx *fiber.Ctx) error {

		imageFile, err := ctx.FormFile("image")

		if err != nil {
			log.Println("image upload error --> ", err)
			return ctx.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}

		ctx.SaveFile(imageFile, "./images/my_image.png")

		return ctx.JSON(fiber.Map{"status": 200, "msg": "Successfully uploaded"})
	})

	app.Listen(":3000")
}
