package service

import (
	"context"

	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

type UserService interface {
	GetByID(ctx context.Context, ID int64) (*user.User, error)
	GetByIDs(ctx context.Context, IDs []int64) ([]*user.User, error)
	GetByNickName(ctx context.Context, nickName string) ([]*user.User, error)
	GetByEmailAndPasswordHash(ctx context.Context, email, passwordHash string) (*user.User, error)
	GetByPhone(ctx context.Context, phone string) ([]*user.User, error)

	Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error)

	Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error)

	DeleteByID(ctx context.Context, ID int64) error
}
