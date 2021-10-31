package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(userProfile dto.UserProfile) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, userProfile.Name, userProfile.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUsers() ([]entity.User, error) {
	var users []entity.User

	query := fmt.Sprintf("SELECT u.id, u.name, u.email FROM %s u", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *AuthPostgres) DeleteUser(userId dto.UserId) error {
	query := fmt.Sprintf("DELETE FROM %s u WHERE u.id = $1", usersTable)
	_, err := r.db.Exec(query, userId.Id)

	return err
}
