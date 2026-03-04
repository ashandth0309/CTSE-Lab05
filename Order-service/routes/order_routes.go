package routes

import (
	"order-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(app *fiber.App) {
	orderGroup := app.Group("/")

	orderGroup.Get("/", controllers.GetOrders)
	orderGroup.Post("/", controllers.CreateOrder)
	orderGroup.Get("/:id", controllers.GetOrderByID)
}
