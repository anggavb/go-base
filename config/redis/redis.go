package config_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New(addr, username, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	cmdErr := rdb.Ping(context.Background())

	return rdb, cmdErr.Err()
}
