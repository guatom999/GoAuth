package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/guatom999/go-auth/service"
)

type catalogHandler struct {
	productService service.ProductServicee
}

func NewCatalogHandler(productService service.ProductServicee) ProductHandler {
	return catalogHandler{productService}
}

func (obj catalogHandler) GetProducts(c *fiber.Ctx) (err error) {

	product, err := obj.productService.GetAllProducts()

	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{
		"products": product,
	})
}
