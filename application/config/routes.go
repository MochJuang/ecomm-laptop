package config

import (
	"github.com/MochJuang/ecomm-laptop/application/controller"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"github.com/MochJuang/ecomm-laptop/application/seeder"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = SetupDatabaseConnection()
	bookRepository  repository.BrandRepository = repository.NewBrandRepository(db)
	brandService    service.BrandService       = service.NewBrandService(bookRepository)
	brandController controller.BrandController = controller.NewBrandController(brandService)
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1")

	// DATA SEEDER
	seeder.BrandSeeder{}.Run(db)

	brand := api.Group("/brand")
	brand.Get("/", brandController.Get)

}
