package config

import (
	"github.com/MochJuang/ecomm-laptop/application/controller"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"github.com/MochJuang/ecomm-laptop/application/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = SetupDatabaseConnection()
	brandRepository     repository.BrandRepository     = repository.NewBrandRepository(db)
	brandService        service.BrandService           = service.NewBrandService(brandRepository)
	brandController     controller.BrandController     = controller.NewBrandController(brandService)
	merkRepository      repository.MerkRepository      = repository.NewMerkRepository(db)
	merkService         service.MerkService            = service.NewMerkService(merkRepository)
	merkController      controller.MerkController      = controller.NewMerkController(merkService)
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	userService         service.UserService            = service.NewUserService(userRepository)
	authController      controller.AuthController      = controller.NewAuthController(userService)
	keranjangRepository repository.KeranjangRepository = repository.NewKeranjangRepository(db)
	keranjangService    service.KeranjangService       = service.NewKeranjangService(keranjangRepository)
	keranjangController controller.KeranjangController = controller.NewKeranjangController(keranjangService)
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1")

	// DATA SEEDER
	// seeder.MainSeeder{}.Run(db)
	// seeder.UserSeeder{}.Run(db)

	brand := api.Group("/brand")
	brand.Get("/", brandController.Get)

	merk := api.Group("/merk")
	merk.Get("/", merkController.Get)
	merk.Get("/:id", merkController.GetDetail)

	auth := api.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)

	keranjang := api.Group("/keranjang")
	keranjang.Post("/", keranjangController.Add)
	keranjang.Get("/:user_id", keranjangController.Get)

}
