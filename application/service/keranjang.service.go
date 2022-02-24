package service

import (
	"log"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
	"github.com/mashingan/smapping"
)

type KeranjangService interface {
	AddKeranjang(k model.KeranjangAddRequest) (bool, error)
	GetKeranjang(userId uint64) ([]model.Keranjang, error)
}

type keranjangService struct {
	keranjangRepository repository.KeranjangRepository
}

func NewKeranjangService(keranjangRepository repository.KeranjangRepository) KeranjangService {
	return &keranjangService{
		keranjangRepository: keranjangRepository,
	}
}

func (ks keranjangService) AddKeranjang(k model.KeranjangAddRequest) (bool, error) {
	var myK model.Keranjang
	err := smapping.FillStruct(&myK, smapping.MapFields(&k))
	myK.StatusKeranjang = "save"
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	_, err = ks.keranjangRepository.Insert(myK)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (ks *keranjangService) GetKeranjang(userId uint64) ([]model.Keranjang, error) {
	return ks.keranjangRepository.Find(userId)
}
