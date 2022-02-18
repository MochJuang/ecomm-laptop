package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BrandRepository interface {
	All() ([]model.Brand, error)
	Insert(brand model.Brand) (model.Brand, error)
}

type brandConnection struct {
	connection *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandConnection{
		connection: db,
	}
}

func (db *brandConnection) Insert(brand model.Brand) (model.Brand, error) {
	result := db.connection.Save(&brand)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return brand, errors.New(result.Error.Error())
	}
	db.connection.Find(&brand)

	return brand, nil
}

func (db *brandConnection) All() ([]model.Brand, error) {
	var brands []model.Brand
	result := db.connection.Find(&brands)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return brands, errors.New(result.Error.Error())
	}
	return brands, nil
}
