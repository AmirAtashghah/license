package api

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) SetRoutes(allowPanel bool, app *fiber.App) {

	api := app.Group("/api")
	api.Post("/license-check", h.CheckLicense)

	api.Post("/token", h.ValidateToken) // deploy change api

	if allowPanel {

		panel := api.Group("/panel", h.GetTokenFromCookie) // deploy chang api

		// user apis :

		// login
		api.Post("/login", h.Login) // deploy chang to api
		// logout
		panel.Get("/logout", h.Logout)

		superAdminPanel := panel.Group("/admin/user", h.SuperAdminCheck)
		// add user
		superAdminPanel.Post("/create", h.CreateUser)
		// update user
		superAdminPanel.Post("/update", h.UpdateUser)
		// get user
		superAdminPanel.Post("/get", h.GetUser)
		// get users
		superAdminPanel.Post("/list", h.ListUsers)
		// delete user
		superAdminPanel.Post("/delete", h.DeleteUser)

		// product apis :
		product := panel.Group("/product")
		// add product
		product.Post("/create", h.CreateProduct)
		// update product
		product.Post("/update", h.UpdateProduct)
		// get product
		product.Post("/get", h.GetProduct)
		// get products
		product.Post("/list", h.ListProducts)
		// delete product
		product.Post("/delete", h.DeleteProduct)

		// customer apis :
		customer := panel.Group("/customer")
		// add customer
		customer.Post("/create", h.CreateCustomer)
		// update customer
		customer.Post("/update", h.UpdateCustomer)
		// get customer
		customer.Post("/get", h.GetCustomer)
		// get customers
		customer.Post("/list", h.ListCustomers)
		// delete customer
		customer.Post("/delete", h.DeleteCustomer)

		// customerProduct apis :
		customerProduct := panel.Group("/customer-product")
		// add customerProduct
		customerProduct.Post("/create", h.CreateCustomerProduct)
		// update customerProduct
		customerProduct.Post("/update", h.UpdateCustomerProduct)
		// get customerProduct
		customerProduct.Post("/get", h.GetCustomerProduct)
		// get customerProduct
		customerProduct.Post("/list", h.ListCustomerProducts)
		// delete customerProduct
		customerProduct.Post("/delete", h.DeleteCustomerProduct)

		// restriction apis :
		restriction := panel.Group("/restriction")
		// add restriction
		restriction.Post("/create", h.CreateRestriction)
		// update restriction
		restriction.Post("/update", h.UpdateRestriction)
		// get restriction
		restriction.Post("/get", h.GetRestriction)
		// get restrictions
		restriction.Post("/list", h.ListRestrictions)
		// delete restriction
		restriction.Post("/delete", h.DeleteRestriction)

		// log apis :
		logs := panel.Group("/logs")
		// get logs license check
		logs.Post("/list-license-check", h.GetLicenseCheckActivityLog)
		// get logs by title
		logs.Post("/list-by-title", h.ListActivityLogsByTitle)
		//  get all logs
		logs.Post("/list", h.ListActivityLogs)

		logs.Get("/count-records", h.CountRecords)
		//deploy un comment this
		app.Static("/", "./static", fiber.Static{
			//ModifyResponse: func(ctx *fiber.Ctx) error {
			//	//ctx.Set("Clear-Site-Data", "\"cache\"")
			//	ctx.Set("Cache-Control", "no-store")
			//	return nil
			//},
			Download: false,
			Compress: true,
		})
		app.Get("/*", func(ctx *fiber.Ctx) error {
			//ctx.Set("Clear-Site-Data", "cache")
			//ctx.Set("Cache-Control", "no-store")
			return ctx.SendFile("./static/index.html")
		})
	}

}
