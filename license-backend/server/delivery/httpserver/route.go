package httpserver

import "github.com/gofiber/fiber/v2"

func (h Handler) SetRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/license/check", h.CheckLicense)

	panel := app.Group("/panel", h.GetTokenFromCookie)

	panel.Post("/login", h.Login)

	// list client api

	// update client api

	// delete client api

	// create client api : create manually

	// show log api

	// deactivate client license

	// activate client license

	// todo setting , add user if need

}
