package auth

import (
	"repositorie/internal/service"
	"repositorie/internal/storage"
)

type Service struct {
	userStorage    storage.UserStorage
	userService    service.UserService
	messageService service.MessageService
}

func NewService(
	userStorage storage.UserStorage,
	userService service.UserService,
	messageService service.MessageService,
) *Service {

	return &Service{
		userStorage:    userStorage,
		userService:    userService,
		messageService: messageService,
	}
}
