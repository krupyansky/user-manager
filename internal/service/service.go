package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
	"github.com/krupyansky/user-manager/internal/repository"
)

type Authorization interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	DeleteUser(userId dto.UserId) error
	GetUsers() ([]entity.User, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, reposRedis *repository.Redis) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User, reposRedis.UsersRedis),
	}
}
