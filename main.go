package main

import (
	"fmt"
	"log"

	"github.com/MochJuang/ecomm-laptop/application/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Static("/", "./public")

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET,PUT,POST,DELETE,OPTIONS",
		ExposeHeaders: "Content-Type,Authorization,Accept",
	}))

	app.Static("/", "./public")

	app.Use(func(c *fiber.Ctx) error {
		log.Println(c.Method() + " " + c.Path())
		return c.Next()
	})

	config.Routes(app)

	fmt.Println("Server started")
	app.Listen(":3500")

}
