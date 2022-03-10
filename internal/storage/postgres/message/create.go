package message

import (
	"context"
	"fmt"
	"repositorie/internal/entities/message"
	"time"
)

func (s *Store) Create(ctx context.Context, q *message.CreateMessageQuery) (*message.Message, error) {
	query := fmt.Sprintf(`
INSERT INTO %s
	(text, chat_id, peer_id)
VALUES
	($1,$2,$3)
RETURNING id
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.Text, q.ChatID, q.PeerID)
	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id after message create: %w", err)
	}

	return &message.Message{
		ID:        id,
		Text:      q.Text,
		ChatID:    q.ChatID,
		PeerID:    q.PeerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
