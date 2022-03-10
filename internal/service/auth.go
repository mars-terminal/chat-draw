package service

import (
	"context"
	"repositorie/internal/entities/auth"
	"repositorie/internal/entities/user"
)

type AuthService interface {
	SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error)
	SignUp(ctx context.Context, q *auth.SignUpQuery) (*auth.Tokens, error)
	GetUserByToken(ctx context.Context, token string) (*user.User, error)
}
