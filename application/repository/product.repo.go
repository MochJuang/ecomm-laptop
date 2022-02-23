package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product model.Product) (model.Product, error)
	// Find(product model.Product) error
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) Insert(product model.Product) (model.Product, error) {
	result := db.connection.Create(&product)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return product, errors.New(result.Error.Error())
	}
	return product, nil
}

// func (db *productConnection) Find(product *model.Product) (error) {

// }
