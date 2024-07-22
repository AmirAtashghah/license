package main

import (
	"github.com/gofiber/fiber/v2"
	api "server/delivery/httpserver"
	"server/repository/postgresql"
)

func init() {

	postgresql.PostgreConnection()
	postgresql.RedisConnection()
}

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Post("/check-license", api.CheckLicense)
	app.Post("/")

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
