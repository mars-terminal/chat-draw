package auth

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

func (s *Service) SignUp(ctx context.Context, q *auth.SignUpQuery) (*auth.Tokens, error) {
	q.Password = s.generatePasswordHash(q.Password)

	u, err := s.userService.Create(ctx, &user.CreateUserQuery{
		FirstName:  q.FirstName,
		SecondName: q.SecondName,
		NickName:   q.NickName,
		Email:      q.Email,
		Phone:      q.Phone,
		Password:   q.Password,
	})
	if err != nil {
		return nil, err
	}

	return s.generateTokens(ctx, u)
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}
