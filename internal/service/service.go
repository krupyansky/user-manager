package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
	"github.com/krupyansky/user-manager/internal/repository"
)

type User interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	DeleteUser(userId dto.UserId) error
	GetUsers() ([]entity.User, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository, reposRedis *repository.Redis) *Service {
	return &Service{
		User: NewUserService(repos.User, reposRedis.UsersRedis),
	}
}
