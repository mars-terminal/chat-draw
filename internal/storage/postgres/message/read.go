package message

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mars-terminal/chat-draw/internal/entities"
	"github.com/mars-terminal/chat-draw/internal/entities/message"
)

func (s *Store) GetByID(ctx context.Context, ID int64) (*message.Message, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE id=$1`, s.table)

	row := s.db.QueryRowxContext(ctx, query, ID)
	if err := row.Err(); err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("message: %w", entities.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to find message by id: %w", err)
	}
	var m message.Message
	err := row.StructScan(&m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Store) GetByChatIDAndPeerID(ctx context.Context, chatID, peerID, limit, offset int64) ([]*message.Message, error) {
	query := fmt.Sprintf(`
SELECT
	id, chat_id, peer_id, text, created_at, updated_at
FROM 
	%s 
WHERE
	(chat_id=$1 AND peer_id=$2) OR (chat_id=$3 AND peer_id=$4)
ORDER BY
	created_at DESC
LIMIT $5 OFFSET $6
`, s.table)

	rows, err := s.db.QueryxContext(ctx, query, chatID, peerID, peerID, chatID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get by chat id (message store): %w", err)
	}

	var messages = make([]*message.Message, 0)
	for rows.Next() {
		var m message.Message
		err := rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func (s *Store) Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error) {
	searchQuery := fmt.Sprintf(`
SELECT 
	text
FROM 
	%s
WHERE text ILIKE $1 AND chat_id = $2
LIMIT ?
OFFSET ?
`, s.table)

	searchQuery = fmt.Sprintf("%%%s%%", searchQuery)

	rows, err := s.db.QueryxContext(ctx, searchQuery, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to search (message store): %w", err)
	}

	var messages []*message.Message
	for rows.Next() {
		var m message.Message
		err := rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}

	return messages, nil
}
