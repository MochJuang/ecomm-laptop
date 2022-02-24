package service

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
)

type UserService interface {
	Login(user model.User) (model.UserToken, error)
	Register(user model.User) (model.UserToken, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Login(user model.User) (model.UserToken, error) {
	res, err := s.userRepository.Login(user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *userService) Register(user model.User) (model.UserToken, error) {
	pass := md5.Sum([]byte(user.Password))
	token := sha256.Sum256([]byte(user.Email))
	res, err := s.userRepository.Insert(model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: fmt.Sprintf("%x", pass),
		Token:    fmt.Sprintf("%x", token),
	})
	if err != nil {
		return res, err
	}
	return res, nil
}
