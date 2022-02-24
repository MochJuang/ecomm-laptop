package controller

import (
	"fmt"
	"log"

	"github.com/MochJuang/ecomm-laptop/application/constants"
	"github.com/MochJuang/ecomm-laptop/application/helper"
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authController struct {
	userService service.UserService
}

func NewAuthController(bs service.UserService) AuthController {
	return &authController{
		userService: bs,
	}
}

var validate *validator.Validate = validator.New()

func (a *authController) Login(c *fiber.Ctx) error {
	var userR model.LoginRequest

	if err := c.BodyParser(&userR); err != nil {
		fmt.Println("Error parsing user")
		fmt.Println(err.Error())
		return helper.BuildErrorResponse(c, constants.ErrorAction, "", helper.EmptyObj{})
	}

	if err := validate.Struct(userR); err != nil {

		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}
	res, err := a.userService.Login(model.User{
		Email:    userR.Email,
		Password: userR.Password,
	})

	if err != nil {
		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}
	return helper.BuildResponse(c, "Success ", res)
}
func (a *authController) Register(c *fiber.Ctx) error {
	var userR model.RegisterRequest

	if err := c.BodyParser(&userR); err != nil {
		fmt.Println("Error parsing user")
		fmt.Println(err.Error())
		return helper.BuildErrorResponse(c, constants.ErrorAction, "", helper.EmptyObj{})
	}

	if err := validate.Struct(userR); err != nil {
		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}
	log.Println(userR)
	if userR.Password != userR.ConfirmPassword {
		return helper.BuildErrorResponse(c, "Error ", "Password not match", helper.EmptyObj{})
	}
	res, err := a.userService.Register(model.User{
		Name:     userR.Name,
		Email:    userR.Email,
		Password: userR.Password,
	})

	if err != nil {
		return helper.BuildErrorResponse(c, "Error ", err.Error(), helper.EmptyObj{})
	}
	return helper.BuildResponse(c, "Success ", res)
}
