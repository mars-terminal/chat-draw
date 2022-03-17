package message

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(map[string]interface{}{
	"package": "message",
	"layer":   "storage",
})

type Store struct {
	db *sqlx.DB

	table string
}

func NewStore(db *sqlx.DB, table string) *Store {
	return &Store{
		db:    db,
		table: table,
	}
}
