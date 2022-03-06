package service

import (
	"context"
	"repositorie/internal/entities/auth"
	"repositorie/internal/entities/user"
)

type AuthService interface {
	SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error)
	SignUp(ctx context.Context, q *auth.SignUpQuery) (*user.User, error)
}
