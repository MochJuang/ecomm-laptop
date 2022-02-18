package seeder

import (
	"sync"

	"github.com/MochJuang/ecomm-laptop/application/config"
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"gorm.io/gorm"
)

type BrandSeeder struct{}

func (bs *BrandSeeder) Run() {
	group := &sync.WaitGroup{}

	var (
		db              *gorm.DB                   = config.SetupDatabaseConnection()
		brandRepository repository.BrandRepository = repository.NewBrandRepository(db)
	)

	listBrand := []map[string]string{
		{
			"brand": "ASUS",
			"image": "",
		},
		{
			"brand": "LENOVO",
			"image": "",
		},
		{
			"brand": "ACER",
			"image": "",
		},
		{
			"brand": "SAMSUNG",
			"image": "",
		},
		{
			"brand": "DELL",
			"image": "",
		},
		{
			"brand": "MAC",
			"image": "",
		},
		{
			"brand": "ZYREX",
			"image": "",
		},
		{
			"brand": "HP",
			"image": "",
		},
	}

	for i := 0; i < len(listBrand); i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			brandRepository.Insert(model.Brand{
				Brand: listBrand[i]["brand"],
				Image: listBrand[i]["image"],
			})
		}()

		group.Wait()
	}
}
