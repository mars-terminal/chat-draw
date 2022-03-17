package service

import (
	"context"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

type AuthService interface {
	SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error)
	SignUp(ctx context.Context, q *auth.SignUpQuery) (*auth.Tokens, error)
	GetUserByToken(ctx context.Context, token string) (*user.User, error)
}
