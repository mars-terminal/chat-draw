package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"repositorie/internal/entities"
)

type MessageStore struct {
	db *sqlx.DB

	table string
}

func NewMessageStore(db *sqlx.DB, table string) *MessageStore {
	return &MessageStore{
		db:    db,
		table: table,
	}
}

func (m *MessageStore) GetByID(ctx context.Context, ID int64) (*entities.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id (message store): %w", err)
	}

	message := entities.Message{}
	for rows.Next() {
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}
	}

	return &message, nil
}

func (m *MessageStore) GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error) {
	query := fmt.Sprintf(`
SELECT
	*
FROM 
	%s 
WHERE
	chat_id=?
LIMIT ? OFFSET ?
`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failet to get by chat id (message store): %w", err)
	}

	var messages []*entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil

}

func (m *MessageStore) GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error) {
	query := fmt.Sprintf(`
SELECT 
	*
FROM 
	%s
WHERE 
	peer_id = ?
LIMIT ?
OFFSET ?`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get by peer id (message store): %w", err)
	}

	var messages []*entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil

}

func (m *MessageStore) Create(ctx context.Context, q *entities.CreateMessageQuery) (*entities.Message, error) {
	query := fmt.Sprintf(`
INSERT INTO %s
	(text, chat_id, peer_id)
VALUES
	(?,?,?)
RETURNING id
`, m.table)

	result, err := m.db.ExecContext(ctx, query, q.Text, q.ChatID, q.PeerID)
	if err != nil {
		return nil, fmt.Errorf("failet to create (message store): %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failet to last insert id in create (message store): %w", err)
	}

	return &entities.Message{
		ID:        id,
		Text:      q.Text,
		ChatID:    q.ChatID,
		PeerID:    q.PeerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *MessageStore) Search(ctx context.Context, query string, limit, offset int64) ([]*entities.Message, error) {
	searchQuery := fmt.Sprintf(`
SELECT 
	*
FROM 
	%s
WHERE text ILIKE '%%?%%'
LIMIT ?
OFFSET ?
`, m.table)

	rows, err := m.db.QueryxContext(ctx, searchQuery, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failet to search (message store): %w", err)
	}

	var messages []*entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.StructScan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

func (m *MessageStore) Update(ctx context.Context, q *entities.UpdateMessageQuery) (*entities.Message, error) {
	query := fmt.Sprintf(`
UPDATE 
	%s
SET text = ?,
	updated_at = now()
WHERE 
	id = ?
`, m.table)

	result, err := m.db.ExecContext(ctx, query, q.Text, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute update %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}

	message, err := m.GetByID(ctx, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id message: %w", err)
	}

	return message, nil

}

func (m *MessageStore) DeleteByID(ctx context.Context, ID int64) error {
	query := fmt.Sprintf(`
DELETE FROM %s
WHERE id=?
`, m.table)

	result, err := m.db.ExecContext(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("failed to delete (message store): %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return nil
	}

	return nil

}
