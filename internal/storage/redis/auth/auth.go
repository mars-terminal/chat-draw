package auth

import (
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	db *redis.Client

	prefix string
}

func NewAuthStorage(db *redis.Client, prefix string) *Storage {
	return &Storage{db: db, prefix: prefix}
}
