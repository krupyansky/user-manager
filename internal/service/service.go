package service

import (
	"github.com/krupyansky/user-manager/internal/entity"
	"github.com/krupyansky/user-manager/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
