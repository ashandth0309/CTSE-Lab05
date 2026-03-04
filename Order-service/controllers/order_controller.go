package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var orders = []fiber.Map{}
var idCounter = 1

func GetOrders(c *fiber.Ctx) error {
	return c.JSON(orders)
}

func CreateOrder(c *fiber.Ctx) error {
	type Order struct {
		Item       string `json:"item"`
		Quantity   int    `json:"quantity"`
		CustomerID string `json:"customerId"`
	}

	var order Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	newOrder := fiber.Map{
		"id":         idCounter,
		"item":       order.Item,
		"quantity":   order.Quantity,
		"customerId": order.CustomerID,
		"status":     "PENDING",
	}

	idCounter++
	orders = append(orders, newOrder)

	return c.Status(201).JSON(newOrder)
}

func GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, order := range orders {
		if id == fmt.Sprintf("%v", order["id"]) {
			return c.JSON(order)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "Order not found",
	})
}
