package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/krupyansky/user-manager/internal/entity"
)

type UsersRedis interface {
	SaveUsers(users []entity.User)
	GetUsers() []entity.User
}

type Redis struct {
	UsersRedis
}

func NewRepositoryRedis(client *redis.Client) *Redis {
	return &Redis{
		UsersRedis: NewUserRedis(client),
	}
}
