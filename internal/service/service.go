package service

import (
	"context"
	"repositorie/internal/entities/auth"
)

type AuthService interface {
	SignIn(ctx context.Context, q *auth.SignInQuery) (*auth.Tokens, error)
	SignUp(ctx context.Context, q *auth.SignUpQuery) (*auth.Tokens, error)
}

type UserService interface {
}

type MessageService interface {
}
