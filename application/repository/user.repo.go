package repository

import (
	"errors"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user model.User) (model.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Insert(user model.User) (model.User, error) {
	result := db.connection.Create(&user)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return user, errors.New(result.Error.Error())
	}
	return user, nil
}
