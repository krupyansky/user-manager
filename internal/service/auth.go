package service

import (
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(userProfile dto.UserProfile) (int, error) {
	return s.repo.CreateUser(userProfile)
}

func (s *AuthService) DeleteUser(userId dto.UserId) error {
	return s.repo.DeleteUser(userId)
}
