package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MemoryRepository interface {
	Insert(memory model.Memory) (model.Memory, error)
}

type memoryConnection struct {
	connection *gorm.DB
}

func NewMemoryRepository(db *gorm.DB) MemoryRepository {
	return &memoryConnection{
		connection: db,
	}
}

func (db *memoryConnection) Insert(memory model.Memory) (model.Memory, error) {
	result := db.connection.Create(&memory)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return memory, errors.New(result.Error.Error())
	}
	return memory, nil
}
