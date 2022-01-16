package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(username, password, host, port, dbname string) (*Store, error) {
	dataSourceName := fmt.Sprintf(
		"username=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		username, password, host, port, dbname,
	)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Store{
		db: db,
	}, nil
}
