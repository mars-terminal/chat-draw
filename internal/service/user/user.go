package user

import (
	"repositorie/internal/storage"
)

type Service struct {
	user storage.UserStorage
}

func NewService(user storage.UserStorage) *Service {
	return &Service{user: user}
}
