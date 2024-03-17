package service

import (
	"Quest/internal/types"
	"Quest/internal/utils"
	"errors"
)

type Repository interface {
	AddUser(User types.User) error
}

type Service struct {
	repo Repository
}

func InitService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddUser(user types.User) error {
	validate := utils.ValidateUser(user)
	if !validate {
		return errors.New("Can't validate user")
	}
	return s.repo.AddUser(user)
}
