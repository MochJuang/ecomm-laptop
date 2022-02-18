package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MerkRepository interface {
	Insert(merk model.Merk) (model.Merk, error)
}

type MerkConnection struct {
	connection *gorm.DB
}

func NewMerkRepository(db *gorm.DB) MerkRepository {
	return &MerkConnection{
		connection: db,
	}
}

func (db *MerkConnection) Insert(merk model.Merk) (model.Merk, error) {
	result := db.connection.Save(&merk)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return merk, errors.New(result.Error.Error())
	}
	db.connection.Find(&merk)

	return merk, nil
}
