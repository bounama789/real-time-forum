package models

import (
	"github.com/gofrs/uuid/v5"
)

type Chat struct {
	ChatId    uuid.UUID `json:"chat_id"`
	CreatedAt string    `json:"created_at"`
}
