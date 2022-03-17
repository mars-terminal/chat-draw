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

func (s *Store) GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error) {
	query := fmt.Sprintf(`
SELECT
	chat_id
FROM 
	%s 
WHERE
	chat_id=$1
LIMIT ? OFFSET ?
`, s.table)

	rows, err := s.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get by chat id (message store): %w", err)
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

func (s *Store) GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*message.Message, error) {
	query := fmt.Sprintf(`
SELECT 
	peer_id
FROM 
	%s
WHERE 
	peer_id=$1
LIMIT ?
OFFSET ?`, s.table)

	rows, err := s.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get by peer id (message store): %w", err)
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

func (s *Store) Search(ctx context.Context, query string, limit, offset int64) ([]*message.Message, error) {
	searchQuery := fmt.Sprintf(`
SELECT 
	text
FROM 
	%s
WHERE text ILIKE '%%?%%'
LIMIT ?
OFFSET ?
`, s.table)

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
