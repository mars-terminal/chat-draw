package user

import (
	"context"
	"fmt"
	"time"

	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

func (s *Store) Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s 
	(first_name, second_name, nick_name, email, phone, password)
VALUES 
	($1, $2, $3, $4, $5, $6)
RETURNING id
`, s.table)

	row := s.db.QueryRowxContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Email, q.Phone, q.Password)

	var id int64
	err := row.StructScan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to last insert id in create (user storage): %w", err)
	}

	return &user.User{
		ID:         id,
		FirstName:  q.FirstName,
		SecondName: q.SecondName,
		NickName:   q.NickName,
		Email:      q.Email,
		Phone:      q.Phone,
		Password:   q.Password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
