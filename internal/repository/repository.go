package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
)

type Authorization interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(userId dto.UserId) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
