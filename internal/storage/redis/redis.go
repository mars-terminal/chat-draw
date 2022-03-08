package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host string
	Port string
}

func NewRedisStorage(ctx context.Context, cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	if status := rdb.Ping(ctx); status.Err() != nil {
		return nil, status.Err()
	}

	return rdb, nil
}
