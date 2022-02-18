package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(c *fiber.Ctx, message string, data interface{}) error {
	res := Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return c.JSON(res)
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(c *fiber.Ctx, message string, err string, data interface{}) error {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return c.Status(fiber.StatusBadRequest).JSON(res)
}
