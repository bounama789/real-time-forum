package models

import (
	"github.com/gofrs/uuid/v5"
)

type Chat struct {
	ChatId    uuid.UUID `json:"chat_id"`
	RequesterId uuid.UUID `json:"requester_id"`
	RecipientId uuid.UUID `json:"recipient_id"`
	CreatedAt string    `json:"created_at"`
	LastMessageTime string `json:"last_message_time"`
}
