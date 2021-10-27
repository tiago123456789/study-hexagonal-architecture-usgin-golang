package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiago123456789/study-hexagonal-architecture-usgin-golang/application"
)

type erroReponse struct {
	Message string `json:"message"`
}

type ProductHandler struct {
	service application.ProductServiceInterface
}

func NewProductHandler(service application.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p *ProductHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := p.service.Get(id)
	if err != nil || product == nil {
		return c.Status(404).JSON(erroReponse{
			Message: "Product not found",
		})
	}

	return c.JSON(product)
}

func (p *ProductHandler) Save(c *fiber.Ctx) error {
	product := new(application.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	productCreated, err := p.service.Create(product.Name, product.Price)
	if err != nil {
		return c.Status(500).JSON(erroReponse{
			Message: "Internal server error",
		})
	}

	return c.JSON(productCreated)
}
