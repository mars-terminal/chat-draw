package user

import (
	"context"
	"fmt"

	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

func (s *Store) Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error) {
	query := fmt.Sprintf(`
UPDATE %s 
SET 
	first_name=$1,
	secont_name = $2,
	nick_name = $3,
	phone = $4,
	password = $5
WHERE 
	id = $6
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failet to update (user storage): %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failet to last insert id in update (user store): %w", err)
	}

	message, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id in update (user store): %w", err)
	}

	return message, nil
}
