package postgres

import (
	"context"
	"fmt"

	"repositorie/internal/entities"
)

type MessageStore struct {
	*Store

	table string
}

func NewMessageStore(store *Store, table string) *MessageStore {
	return &MessageStore{
		Store: store,
		table: table,
	}
}

func (m *MessageStore) GetByID(ctx context.Context, ID int64) (*entities.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=? LIMIT 1`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID)
	if err != nil {
		return nil, err
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
	query := fmt.Sprintf(`SELECT * FROM %s WHERE chat_id in (?) LIMIT ? OFFSET ?`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, err
	}

	var messages = make([]*entities.Message, 0)
	for rows.Next() {
		err := rows.StructScan(&messages)
		if err != nil {
			return nil, err
		}
	}

	return messages, nil
}

func (m *MessageStore) GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE peer_id in (?) LIMIT ? OFFSET ?`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID, limit, offset)
	if err != nil {
		return nil, err
	}

	var messages = make([]*entities.Message, 0)
	for rows.Next() {
		err := rows.StructScan(&messages)
		if err != nil {
			return nil, err
		}
	}

	return messages, nil
}

func (m *MessageStore) Create(ctx context.Context, q *entities.CreateMessageQuery) (*entities.Message, error) {
	query := fmt.Sprintf(`
INSERT INTO %s
	(text, chat_id, peer_id)
VALUES
	(?, ?, ?)
`, m.table)

	result, err := m.db.ExecContext(ctx, query, q.Text, q.ChatID, q.PeerID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	message, err := m.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m *MessageStore) Search(ctx context.Context, q string, limit, offset int64) ([]*entities.Message, error) {
	query := fmt.Sprintf(`
SELECT *
FROM %s
WHERE text ILIKE '%%?%%'
`, m.table)

	rows, err := m.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	messages := make([]*entities.Message, 0)
	for rows.Next() {
		message := entities.Message{}

		if err := rows.StructScan(&message); err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (m *MessageStore) Update(ctx context.Context, q *entities.UpdateMessageQuery) (*entities.Message, error) {
	query := fmt.Sprintf(`
INSERT INTO %s
	(text)
VALUES 
	(?)
RETURNING id
`, m.table)

	result, err := m.db.ExecContext(ctx, query, q.Text)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	message, err := m.GetByID(ctx, id)
	if err != nil {
		return nil, err
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
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found")
	}

	return nil

}
