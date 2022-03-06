package message

import (
	"repositorie/internal/storage/postgres/message"
)

type Service struct {
	message message.Store
}

func NewMessageService(message message.Store) *Service {
	return &Service{
		message: message,
	}
}
