package storage

import (
	"context"
	"repositorie/internal/entities"
)

type UserStorage interface {
	GetByID(ctx context.Context, ID int64) (*entities.User, error)
	GetByIDs(ctx context.Context, IDs []int64) ([]*entities.User, error)
	GetByNickName(ctx context.Context, nickName string) ([]*entities.User, error)
	GetByPhone(ctx context.Context, phone string) ([]*entities.User, error)

	Create(ctx context.Context, q *entities.CreateUserQuery) (*entities.User, error)

	Update(ctx context.Context, q *entities.UpdateUserQuery) (*entities.User, error)

	DeleteByID(ctx context.Context, ID int64) error
}

type MessageStorage interface {
	GetByID(ctx context.Context, ID int64) (*entities.Message, error)

	GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error)
	GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error)

	Create(ctx context.Context, q *entities.CreateMessageQuery) ([]*entities.Message, error)

	Search(ctx context.Context, query string, limit, offset int64) ([]*entities.Message, error)

	Update(ctx context.Context, q *entities.UpdateMessageQuery) ([]*entities.Message, error)

	DeleteByID(ctx context.Context, ID int64) error
}

type Storage struct {
	Users    UserStorage
	Messages MessageStorage
}

func NewStorage(
	users UserStorage,
	messages MessageStorage,
) *Storage {
	return &Storage{
		Users:    users,
		Messages: messages,
	}
}
