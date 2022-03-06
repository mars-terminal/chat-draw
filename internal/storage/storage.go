package storage

import (
	"context"
	"repositorie/internal/entities"
	"repositorie/internal/entities/user"
)

type UserStorage interface {
	GetByID(ctx context.Context, ID int64) (*user.User, error)
	GetByIDs(ctx context.Context, IDs []int64) ([]*user.User, error)
	GetByNickName(ctx context.Context, nickName string) ([]*user.User, error)
	GetByPhone(ctx context.Context, phone string) ([]*user.User, error)

	Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error)

	Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error)

	DeleteByID(ctx context.Context, ID int64) error
}

type MessageStorage interface {
	GetByID(ctx context.Context, ID int64) (*entities.Message, error)

	GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error)
	GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error)

	Create(ctx context.Context, q *entities.CreateMessageQuery) (*entities.Message, error)

	Search(ctx context.Context, query string, limit, offset int64) ([]*entities.Message, error)

	Update(ctx context.Context, q *entities.UpdateMessageQuery) (*entities.Message, error)

	DeleteByID(ctx context.Context, ID int64) error
}
