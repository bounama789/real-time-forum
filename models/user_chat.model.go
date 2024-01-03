package models

import (
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/websocket"
)

type UsersChats struct {
	Conn   *websocket.Conn `json:"conn"`
	UserId string          `json:"user_id"`
	ChatId uuid.UUID       `json:"chat_id"`
}
