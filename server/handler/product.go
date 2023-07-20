package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetTotalProductsPrice(c *fiber.Ctx) error
	// GetById(id int) (*product, error)
}
