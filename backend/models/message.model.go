package models

import "github.com/gofrs/uuid/v5"

type Message struct {
	MessId   uuid.UUID `json:"mess_id"`
	ChatId   uuid.UUID `json:"chat_id"`
	SenderId uuid.UUID `json:"sender_id"`
	Body     string    `json:"body"`
}
