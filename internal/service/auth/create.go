package auth

import (
	"context"
	"repositorie/internal/entities/auth"
	"repositorie/internal/entities/user"
)

func (s *Service) SignUp(ctx context.Context, q *auth.SignUpQuery) (*user.User, error) {
	return s.userStorage.Create(ctx, &user.CreateUserQuery{
		FirstName:  q.FirstName,
		SecondName: q.SecondName,
		NickName:   q.NickName,
		Phone:      q.Phone,
		Password:   q.Password,
	})

}
