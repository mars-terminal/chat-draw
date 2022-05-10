package message

import (
	"context"

	"github.com/mars-terminal/chat-draw/internal/entities/message"
	"github.com/mars-terminal/chat-draw/internal/storage"
)

type Service struct {
	storage storage.MessageStorage
}

func NewService(storage storage.MessageStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetByID(ctx context.Context, ID int64) (*message.Message, error) {
	return s.storage.GetByID(ctx, ID)
}

func (s *Service) GetByChatIDAndPeerID(ctx context.Context, chatID, peerID, limit, offset int64) ([]*message.Message, error) {
	return s.storage.GetByChatIDAndPeerID(ctx, chatID, peerID, limit, offset)
}

func (s *Service) Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error) {
	return s.storage.Create(ctx, q)
}

func (s *Service) Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error) {
	return s.storage.Search(ctx, query, limit, offset)
}

func (s *Service) Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error) {
	return s.storage.Update(ctx, q)
}

func (s *Service) DeleteByID(ctx context.Context, ID int64) error {
	return s.storage.DeleteByID(ctx, ID)
}
