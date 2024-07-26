package httpserver

import "github.com/gofiber/fiber/v2"

func (h Handler) SetRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/license-check/:hash", h.CheckLicense)

	panel := app.Group("/panel", h.GetTokenFromCookie)

	panel.Post("/login", h.Login)

	client := panel.Group("/client")

	// list client api
	client.Post("/list", h.ListClients)

	// update client api
	client.Post("/update", h.UpdateClient)

	// delete client api
	client.Post("/delete", h.DeleteClient)

	// change activate status client license
	client.Post("/change-activate-status", h.UpdateActivateStatus)

	client.Post("/get", h.GetClient)

	// create client api : create manually todo add if needed

	// show log api
	client.Post("/logs/get", h.ListLogs)

	// todo setting , add user if need

}
