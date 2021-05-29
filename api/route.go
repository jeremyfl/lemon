package api

import (
	"customer/api/controller"
	"github.com/gofiber/fiber/v2"
)

// RegisterRoute setting the route of API Controller
// and return as Fiber App
func RegisterRoute(app *fiber.App, controller *controller.Controller) *fiber.App {
	app.Get("/api/v1/customers", controller.CustomerController.Get)

	return app
}
