package entities

import "time"

type Message struct {
	ID        int64     `db:"id"`
	Text      string    `db:"text"`
	ChatID    int64     `db:"chat_id"`
	PeerID    int64     `db:"peer_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateMessageQuery struct {
	Text   string
	ChatID int64
	PeerID int64
}

type UpdateMessageQuery struct {
	ID   int64
	Text string
}
