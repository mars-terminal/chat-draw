package auth

import (
	"context"
	"repositorie/internal/entities/auth"
)

func (s *Service) SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error) {
	u, err := s.userService.GetByNickNameAndPasswordHash(ctx, q.NickName, s.generatePasswordHash(q.Password))
	if err != nil {
		return nil, err
	}

	return s.generateTokens(ctx, u)
}
