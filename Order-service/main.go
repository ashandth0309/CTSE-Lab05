package main

import (
	"log"
	"order-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupOrderRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Order Service Running",
		})
	})

	log.Fatal(app.Listen(":8082"))
}
