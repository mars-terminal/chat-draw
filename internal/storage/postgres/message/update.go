package message

import (
	"context"
	"fmt"
	"repositorie/internal/entities/message"
)

func (s *Store) Update(ctx context.Context, q *message.UpdateMessageQuery) (*message.Message, error) {
	query := fmt.Sprintf(`
UPDATE 
	%s
SET text = $1,
	updated_at = now()
WHERE 
	id = $2
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.Text, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute update %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}

	m, err := s.GetByID(ctx, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id message: %w", err)
	}

	return m, nil

}
