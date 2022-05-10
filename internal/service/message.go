package service

import (
	"context"

	"github.com/mars-terminal/chat-draw/internal/entities/message"
)

type MessageService interface {
	GetByID(ctx context.Context, ID int64) (*message.Message, error)

	GetByChatIDAndPeerID(ctx context.Context, chatID, peerID, limit, offset int64) ([]*message.Message, error)

	Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error)

	Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error)

	Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error)

	DeleteByID(ctx context.Context, ID int64) error
}
