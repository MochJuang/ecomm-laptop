package service

import (
	"log"

	"github.com/MochJuang/ecomm-laptop/application/helper"
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
)

type MerkService interface {
	// Get(prd *model.Merk) error
	GetPaginate(filter model.FilterMerk, page int, limit int) helper.Paginate
	Detail(id uint64) model.DetailMerk
}

type merkService struct {
	merkRepository repository.MerkRepository
}

func NewMerkService(merkRepository repository.MerkRepository) MerkService {
	return &merkService{
		merkRepository: merkRepository,
	}
}

func (pr *merkService) GetPaginate(filter model.FilterMerk, page int, limit int) helper.Paginate {

	if limit == 0 {
		limit = 20
	}

	if page == 0 {
		page = 1
	}

	var offset int
	if page != 0 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}

	totalData, err := pr.merkRepository.FindAndCount(filter)
	if err != nil {
		log.Println(err.Error())
	}
	items, err := pr.merkRepository.FindAndPaginate(filter, limit, offset)
	if err != nil {
		log.Println(err.Error())
	}
	return helper.PaginationHelper(items, page, limit, totalData)
}
func (pr *merkService) Detail(id uint64) model.DetailMerk {
	res, _ := pr.merkRepository.DetailMerk(id)
	return res
}
