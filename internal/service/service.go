package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/repository"
)

type Authorization interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	DeleteUser(userId dto.UserId) error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
