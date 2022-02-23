package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	Get(c *fiber.Ctx) error
}

// type productController struct {
// 	productService service.ProductService
// }
