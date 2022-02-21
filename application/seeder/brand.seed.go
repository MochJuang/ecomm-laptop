package seeder

import (
	"strconv"
	"sync"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"gorm.io/gorm"
)

type BrandSeeder struct {
}

func (bs BrandSeeder) Run(db *gorm.DB) {
	group := &sync.WaitGroup{}

	var (
		brandRepository  repository.BrandRepository  = repository.NewBrandRepository(db)
		memoryRepository repository.MemoryRepository = repository.NewMemoryRepository(db)
		warnaRepository  repository.WarnaRepository  = repository.NewWarnaRepository(db)
	)

	listBrand := []map[string]string{
		{
			"brand": "ASUS",
			"image": "ASUS.png",
		},
		{
			"brand": "LENOVO",
			"image": "LENOVO.png",
		},
		{
			"brand": "ACER",
			"image": "ACER.png",
		},
		{
			"brand": "SAMSUNG",
			"image": "SAMSUNG.png",
		},
		{
			"brand": "DELL",
			"image": "DELL.png",
		},
		{
			"brand": "MAC",
			"image": "MAC.png",
		},
		{
			"brand": "ZYREX",
			"image": "ZYREX.png",
		},
		{
			"brand": "HP",
			"image": "HP.png",
		},
	}
	listWarna := []model.Warna{
		{ID: 1, Warna: "RED"},
		{ID: 2, Warna: "GREEN"},
		{ID: 3, Warna: "BLUE"},
	}

	go func(lists []model.Warna) {
		group.Add(len(listWarna))
		defer group.Done()
		for i := 0; i < len(listWarna); i++ {
			warnaRepository.Insert(listWarna[i])
		}
	}(listWarna)

	listMemory := []model.Memory{
		{ID: 1, Disk: "128 GB SSD", Ram: 4, IsSsd: true},
		{ID: 2, Disk: "256 GB SSD", Ram: 4, IsSsd: true},
		{ID: 3, Disk: "512 GB SSD", Ram: 8, IsSsd: true},
		{ID: 4, Disk: "500 GB SSD", Ram: 16, IsSsd: false},
		{ID: 5, Disk: "1000 GB SSD", Ram: 8, IsSsd: false},
	}

	go func(lists []model.Memory) {
		group.Add(len(listMemory))
		defer group.Done()
		for i := 0; i < len(listMemory); i++ {
			memoryRepository.Insert(listMemory[i])
		}
	}(listMemory)

	listProduct := []*model.Product{
		{
			MemoryId: 1,
			WarnaId:  1,
			Stock:    20,
		},
		{
			MemoryId: 2,
			WarnaId:  1,
			Stock:    20,
		},
		{
			MemoryId: 3,
			WarnaId:  1,
			Stock:    20,
		},
		{
			MemoryId: 4,
			WarnaId:  1,
			Stock:    20,
		},
		{
			MemoryId: 5,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 1,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 2,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 3,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 4,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 5,
			WarnaId:  2,
			Stock:    20,
		},
		{
			MemoryId: 5,
			WarnaId:  3,
			Stock:    20,
		},
		{
			MemoryId: 1,
			WarnaId:  3,
			Stock:    20,
		},
		{
			MemoryId: 2,
			WarnaId:  3,
			Stock:    20,
		},
		{
			MemoryId: 3,
			WarnaId:  3,
			Stock:    20,
		},
		{
			MemoryId: 4,
			WarnaId:  3,
			Stock:    20,
		},
		{
			MemoryId: 5,
			WarnaId:  2,
			Stock:    20,
		},
	}

	go func() {
		group.Add(len(listBrand))
		defer group.Done()
		for i := 0; i < len(listBrand); i++ {
			brandRepository.Insert(model.Brand{
				Brand: listBrand[i]["brand"],
				Image: listBrand[i]["image"],
				Merk: []*model.Merk{
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+1),
						Product: []*model.Product{},
					},
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+2),
						Product: []*model.Product{},
					},
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+3),
						Product: []*model.Product{},
					},
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+4),
						Product: []*model.Product{},
					},
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+5),
						Product: []*model.Product{},
					},
					{
						Merk:    "testing " + listBrand[i]["brand"] + " " + strconv.Itoa(i+6),
						Product: []*model.Product{},
					},
				},
			})

		}
	}()
	group.Wait()
}
