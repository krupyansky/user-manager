package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
	"github.com/krupyansky/user-manager/internal/repository"
)

type AuthService struct {
	repo      repository.User
	repoRedis repository.UsersRedis
}

func NewAuthService(repo repository.User, repoRedis repository.UsersRedis) *AuthService {
	return &AuthService{repo: repo, repoRedis: repoRedis}
}

func (s *AuthService) CreateUser(userProfile dto.UserProfile) (int, error) {
	return s.repo.CreateUser(userProfile)
}

func (s *AuthService) DeleteUser(userId dto.UserId) error {
	return s.repo.DeleteUser(userId)
}

func (s *AuthService) GetUsers() ([]entity.User, error) {
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
