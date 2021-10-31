package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const (
	usersKey     = "users"
	cacheExpires = 60
)

type ConfigRedis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisDB(cfg ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
