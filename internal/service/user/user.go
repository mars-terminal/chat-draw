package user

import (
	"context"
	"errors"
	"fmt"
	"repositorie/internal/entities"
	"repositorie/internal/entities/user"
	"repositorie/internal/storage"
)

type Service struct {
	storage storage.UserStorage
}

func NewService(storage storage.UserStorage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetByID(ctx context.Context, ID int64) (*user.User, error) {
	return s.storage.GetByID(ctx, ID)
}

func (s *Service) GetByIDs(ctx context.Context, IDs []int64) ([]*user.User, error) {
	return s.storage.GetByIDs(ctx, IDs)
}

func (s *Service) GetByNickName(ctx context.Context, nickName string) ([]*user.User, error) {
	return s.storage.GetByNickName(ctx, nickName)
}

func (s *Service) GetByNickNameAndPasswordHash(ctx context.Context, nickName, passwordHash string) (*user.User, error) {
	return s.storage.GetByNickNameAndPasswordHash(ctx, nickName, passwordHash)
}

func (s *Service) GetByPhone(ctx context.Context, phone string) ([]*user.User, error) {
	return s.storage.GetByPhone(ctx, phone)
}

func (s *Service) Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error) {
	_, err := s.storage.GetByNickNameStrict(ctx, q.NickName)
	if err == nil {
		return nil, fmt.Errorf("user already exists: %w", entities.ErrForbidden)
	}

	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return nil, err
	}

	return s.storage.Create(ctx, q)
}

func (s *Service) Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error) {
	return s.storage.Update(ctx, q)
}

func (s *Service) DeleteByID(ctx context.Context, ID int64) error {
	return s.storage.DeleteByID(ctx, ID)
}
