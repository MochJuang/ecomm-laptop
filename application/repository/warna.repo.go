package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarnaRepository interface {
	Insert(warna model.Warna) (model.Warna, error)
}

type warnaConnection struct {
	connection *gorm.DB
}

func NewWarnaRepository(db *gorm.DB) WarnaRepository {
	return &warnaConnection{
		connection: db,
	}
}

func (db *warnaConnection) Insert(warna model.Warna) (model.Warna, error) {
	result := db.connection.Create(&warna)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return warna, errors.New(result.Error.Error())
	}
	return warna, nil
}
