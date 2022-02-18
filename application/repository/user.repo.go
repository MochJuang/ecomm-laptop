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
	result := db.connection.Save(&user)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return user, errors.New(result.Error.Error())
	}
	db.connection.Find(&user)

	return user, nil
}
