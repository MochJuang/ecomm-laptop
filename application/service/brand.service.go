package service

import (
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
)

type BrandService interface {
	GetAll() ([]model.Brand, error)
}

type brandService struct {
	brandRepository repository.BrandRepository
}

func NewBrandService(brandRepository repository.BrandRepository) BrandService {
	return &brandService{
		brandRepository: brandRepository,
	}
}

func (bs *brandService) GetAll() ([]model.Brand, error) {
	res, err := bs.brandRepository.All()
	return res, err
}
