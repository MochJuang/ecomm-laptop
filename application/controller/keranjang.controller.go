package controller

import (
	"fmt"
	"strconv"

	"github.com/MochJuang/ecomm-laptop/application/constants"
	"github.com/MochJuang/ecomm-laptop/application/helper"
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/gofiber/fiber/v2"
)

type KeranjangController interface {
	Add(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
}

type keranjangController struct {
	keranjangService service.KeranjangService
}

func NewKeranjangController(bs service.KeranjangService) KeranjangController {
	return &keranjangController{
		keranjangService: bs,
	}
}
func (kc *keranjangController) Add(c *fiber.Ctx) error {
	var kr model.KeranjangAddRequest

	if err := c.BodyParser(&kr); err != nil {
		fmt.Println("Error parsing user")
		fmt.Println(err.Error())
		return helper.BuildErrorResponse(c, constants.ErrorAction, "", helper.EmptyObj{})
	}

	if err := validate.Struct(kr); err != nil {
		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}

	res, err := kc.keranjangService.AddKeranjang(kr)

	if err != nil {
		return helper.BuildErrorResponse(c, constants.ErrorAction, "", helper.EmptyObj{})
	}

	return helper.BuildResponse(c, "Success ", res)
}
func (kc *keranjangController) Get(c *fiber.Ctx) error {

	idUser, _ := strconv.ParseUint(c.Params("user_id"), 10, 64)
	if idUser == 0 {
		return helper.BuildErrorResponse(c, "Error Validation", "Id User is required ", helper.EmptyObj{})
	}
	res, err := kc.keranjangService.GetKeranjang(idUser)
	if err != nil {
		fmt.Println(err.Error())
		return helper.BuildErrorResponse(c, constants.ErrorAction, "", helper.EmptyObj{})
	}
	return helper.BuildResponse(c, "Success ", res)
}
