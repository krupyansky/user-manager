package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
	"github.com/krupyansky/user-manager/internal/repository"
)

type UserService struct {
	repo      repository.User
	repoRedis repository.UsersRedis
}

func NewUserService(repo repository.User, repoRedis repository.UsersRedis) *UserService {
	return &UserService{repo: repo, repoRedis: repoRedis}
}

func (s *UserService) CreateUser(userProfile dto.UserProfile) (int, error) {
	return s.repo.CreateUser(userProfile)
}

func (s *UserService) DeleteUser(userId dto.UserId) error {
	return s.repo.DeleteUser(userId)
}

func (s *UserService) GetUsers() ([]entity.User, error) {
	users := s.repoRedis.GetUsers()
	if users != nil {
		return users, nil
	}

	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	s.repoRedis.SaveUsers(users)

	return users, nil
}
