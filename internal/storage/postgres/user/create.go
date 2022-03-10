package user

import (
	"context"
	"fmt"
	"repositorie/internal/entities/user"
	"time"
)

func (s *Store) Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s 
	(first_name, second_name, nick_name, phone, password)
VALUES 
	($1, $2, $3, $4, $5)
RETURNING id
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create (user storage): %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to last insert id in create (user storage): %w", err)
	}

	return &user.User{
		ID:         id,
		FirstName:  q.FirstName,
		SecondName: q.SecondName,
		NickName:   q.NickName,
		Phone:      q.Phone,
		Password:   q.Password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
