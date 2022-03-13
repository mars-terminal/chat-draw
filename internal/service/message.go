package service

import (
	"context"
	"repositorie/internal/entities/message"
)

type MessageService interface {
	GetByID(ctx context.Context, ID int64) (*message.Message, error)

	GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error)
	GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error)

	Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error)

	Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error)

	Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error)

	DeleteByID(ctx context.Context, ID int64) error
}