package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/krupyansky/user-manager/internal/entity"
	"time"
)

type UserRedis struct {
	client *redis.Client
}

func NewUserRedis(client *redis.Client) *UserRedis {
	return &UserRedis{client: client}
}

func (r *UserRedis) SaveUsers(users []entity.User) {
	serialized, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	r.client.Set(context.Background(), usersKey, serialized, cacheExpires*time.Second)
}

func (r *UserRedis) GetUsers() []entity.User {
	data, err := r.client.Get(context.Background(), usersKey).Result()
	if err != nil {
		return nil
	}

	var users []entity.User

	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		panic(err)
	}

	return users
}
