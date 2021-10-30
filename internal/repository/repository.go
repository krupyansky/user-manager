package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krupyansky/user-manager/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(userId int) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
