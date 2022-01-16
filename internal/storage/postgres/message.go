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
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, m.table)

	rows, err := m.db.QueryxContext(ctx, query, ID)
	if err != nil {
		return nil, err
	}

	user := entities.User{}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (m *MessageStore) GetByChatID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageStore) GetByPeerID(ctx context.Context, ID int64, limit, offset int64) ([]*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageStore) Create(ctx context.Context, q *entities.CreateMessageQuery) ([]*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageStore) Search(ctx context.Context, query string, limit, offset int64) ([]*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageStore) Update(ctx context.Context, q *entities.UpdateMessageQuery) ([]*entities.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageStore) DeleteByID(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}
