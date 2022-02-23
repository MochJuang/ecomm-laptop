package service

import (
	"github.com/MochJuang/ecomm-laptop/application/model"
	"github.com/MochJuang/ecomm-laptop/application/repository"
)

type UserService interface {
	Login(user model.User) (model.User, error)
	Register(user model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Login(user model.User) (model.User, error) {
	res, err := s.userRepository.Login(user)
	if err != nil {
		return user, err
	}
	return res, nil
}

func (s *userService) Register(user model.User) (model.User, error) {
	res, err := s.userRepository.Insert(user)
	if err != nil {
		return user, err
	}
	return res, nil
}
