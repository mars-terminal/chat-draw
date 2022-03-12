package message

import "time"

type Message struct {
	ID        int64     `json:"id" db:"id"`
	Text      string    `json:"text" db:"text"`
	ChatID    int64     `json:"chat_id" db:"chat_id"`
	PeerID    int64     `json:"peer_id" db:"peer_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CreateMessageQuery struct {
	Text   string `json:"text" validate:"min=1 max=999"`
	ChatID int64  `json:"chat_id"`
	PeerID int64  `json:"peer_id"`
}

type UpdateMessageQuery struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}
