package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
)

type User interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(userId dto.UserId) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
