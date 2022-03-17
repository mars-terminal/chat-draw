package auth

import (
	"context"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

func (s *Service) SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error) {
	u, err := s.userService.GetByNickNameAndPasswordHash(ctx, q.NickName, s.generatePasswordHash(q.Password))
	if err != nil {
		return nil, err
	}

	return s.generateTokens(ctx, u)
}

func (s *Service) GetUserByToken(ctx context.Context, token string) (*user.User, error) {
	userID, err := s.authStorage.GetUserIDByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return s.userService.GetByID(ctx, userID)
}
