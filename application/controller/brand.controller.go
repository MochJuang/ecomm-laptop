package controller

import (
	"github.com/MochJuang/ecomm-laptop/application/helper"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BrandController interface {
	Get(c *fiber.Ctx) error
}

type brandController struct {
	brandService service.BrandService
}

func NewBrandController(bs service.BrandService) BrandController {
	return &brandController{
		brandService: bs,
	}
}

func (b *brandController) Get(c *fiber.Ctx) error {
	brands, err := b.brandService.GetAll()
	if err != nil {
		logrus.Info(err.Error())
		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}
	return helper.BuildResponse(c, "Error ", brands)
}
