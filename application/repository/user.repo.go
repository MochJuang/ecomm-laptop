package repository

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user model.User) (model.UserToken, error)
	Login(user model.User) (model.UserToken, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Insert(user model.User) (model.UserToken, error) {
	result := db.connection.Create(&user)
	if result.Error != nil {
		logrus.Info(result.Error.Error())
		return model.UserToken{}, errors.New(result.Error.Error())
	}
	return model.UserToken{
		Token: user.Token,
		Name:  user.Name,
	}, nil
}

func (db *userConnection) Login(user model.User) (model.UserToken, error) {
	var myUser model.User
	pass := md5.Sum([]byte(user.Password))
	res := db.connection.Where("email = ? AND password = ?", user.Email, fmt.Sprintf("%x", pass)).Find(&myUser)
	if res.RowsAffected == 0 {
		return model.UserToken{}, errors.New("User Not Found")
	}

	return model.UserToken{
		Token: myUser.Token,
		Name:  myUser.Name,
	}, nil
}
