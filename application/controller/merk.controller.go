package controller

import (
	"strconv"

	"github.com/MochJuang/ecomm-laptop/application/helper"
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/gofiber/fiber/v2"
)

type MerkController interface {
	Get(c *fiber.Ctx) error
	GetDetail(c *fiber.Ctx) error
}

type merkController struct {
	merkService service.MerkService
}

func NewMerkController(ms service.MerkService) MerkController {
	return &merkController{
		merkService: ms,
	}
}

func (mc *merkController) Get(c *fiber.Ctx) error {
	brandId, _ := strconv.ParseUint(c.Query("brand_id"), 10, 64)
	ram, _ := strconv.Atoi(c.Query("ram"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	filter := model.FilterMerk{
		Search:  c.Query("search"),
		BrandId: brandId,
		Disk:    c.Query("disk"),
		Ram:     ram,
	}

	result := mc.merkService.GetPaginate(filter, page, limit)
	return helper.BuildResponse(c, "Success ", result)
}

func (mc *merkController) GetDetail(c *fiber.Ctx) error {

	idMerk, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	if idMerk == 0 {
		return helper.BuildErrorResponse(c, "Error ", "Id Merk is required ", helper.EmptyObj{})
	}

	return helper.BuildResponse(c, "Success ", mc.merkService.Detail(idMerk))
}
