package auth

import (
	"repositorie/internal/storage"
)

type Service struct {
	userStorage storage.UserStorage
}

func NewService(userStorage storage.UserStorage) *Service {
	return &Service{userStorage: userStorage}
}
