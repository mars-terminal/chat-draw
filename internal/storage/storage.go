package storage

import (
	"context"
	"time"

	"github.com/mars-terminal/chat-draw/internal/entities/message"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

type AuthStorage interface {
	GetUserIDByToken(ctx context.Context, token string) (int64, error)

	SetToken(ctx context.Context, token string, userID int64, ttl time.Duration) error

	DeleteToken(ctx context.Context, token string) error
}

type UserStorage interface {
	GetByID(ctx context.Context, ID int64) (*user.User, error)
	GetByIDs(ctx context.Context, IDs []int64) ([]*user.User, error)
	GetByNickName(ctx context.Context, nickName string) ([]*user.User, error)
	GetByNickNameStrict(ctx context.Context, nickName string) (*user.User, error)
	GetByNickNameAndPasswordHash(ctx context.Context, nickName, passwordHash string) (*user.User, error)
	GetByPhone(ctx context.Context, phone string) ([]*user.User, error)

	Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error)

	Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error)

	DeleteByID(ctx context.Context, ID int64) error
}

type MessageStorage interface {
	GetByID(ctx context.Context, ID int64) (*message.Message, error)

	GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error)
	GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error)

	Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error)

	Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error)

	Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error)

	DeleteByID(ctx context.Context, ID int64) error
}
