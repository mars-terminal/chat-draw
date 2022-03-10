package message

import (
	"context"
	"repositorie/internal/entities/message"
	"repositorie/internal/storage"
)

type Service struct {
	storage storage.MessageStorage
}

func NewService(storage storage.MessageStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (m *Service) GetByID(ctx context.Context, ID int64) (*message.Message, error) {
	return m.storage.GetByID(ctx, ID)
}

func (m *Service) GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error) {
	return m.storage.GetByChatID(ctx, ID, limit, offset)

}

func (m *Service) GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error) {
	return m.storage.GetByPeerID(ctx, ID, limit, offset)
}

func (m *Service) Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error) {
	return m.Create(ctx, q)
}

func (m *Service) Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error) {
	return m.Search(ctx, query, limit, offset)
}

func (m *Service) Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error) {
	return m.storage.Update(ctx, q)
}

func (m *Service) DeleteByID(ctx context.Context, ID int64) error {
	return m.storage.DeleteByID(ctx, ID)
}
