package main

import (
	"github.com/gofiber/fiber/v2"
	"server/db"
	api "server/delivery/http"
)

func init() {

	db.PostgreConnection()
	db.RedisConnection()
}

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Post("/check-license", api.CheckLicense)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
