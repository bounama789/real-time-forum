package ws

import (
	"encoding/json"
	"forum/backend/models"
	repo "forum/backend/server/repositories"
	"log"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

const (
	WS_JOIN_EVENT       = "join-event"
	WS_DISCONNECT_EVENT = "disconnect-event"
	WS_MESSAGE_EVENT    = "msg-event"
)

type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
}

type WSClient struct {
	Username    string
	WSCoon      *websocket.Conn
	OutgoingMsg chan interface{}
}

type WSPaylaod struct {
	From string
	Type string
	Data interface{}
}

var WSHub *Hub

func init() {
	WSHub = newHub()
	go WSHub.listen()
}

func newHub() *Hub {
	return &Hub{
		Clients:           &sync.Map{},
		RegisterChannel:   make(chan *WSClient),
		UnRegisterChannel: make(chan *WSClient),
	}
}

func (wsHub *Hub) listen() {
	for {
		select {
		case client := <-wsHub.RegisterChannel:
			wsHub.Clients.Store(client.Username, client)
			log.Printf("%s is connected\n", client.Username)
		case client := <-wsHub.UnRegisterChannel:
			wsHub.Clients.Delete(client.Username)
			log.Printf("%s is disconnected\n", client.Username)
		}
	}
}

func (wsHub *Hub) AddClient(coon *websocket.Conn, username string) {
	client := &WSClient{
		Username:    username,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}

	go client.messageReader()
	go client.messageWriter()

	wsHub.RegisterChannel <- client

	var newEvent = WSPaylaod{
		From: client.Username,
		Type: WS_JOIN_EVENT,
		Data: nil,
	}

	wsHub.HandleEvent(newEvent)

}

func (wsHub *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		wsHub.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Username != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGE_EVENT:
		data := eventPayload.Data.(map[string]any)
		c, ok := WSHub.Clients.Load(data["to"])
		client := c.(*WSClient)

		if c, err := repo.ChatRepo.GetChat(client.Username,data["to"].(string)); err != nil || c.ChatId.String() != data["chatId"] {
			return
		}
		if ok {
			var message = models.Message{
				Sender:    client.Username,
				Body:      data["content"].(string),
				CreatedAt: data["time"].(string),
				ChatId:    uuid.FromStringOrNil(data["chatId"].(string)),
			}
			repo.MessRepo.SaveMessage(&message)

			var event = WSPaylaod{
				Type: WS_MESSAGE_EVENT,
				From: client.Username,
				Data: message,
			}

			client.OutgoingMsg <- event
		}
	}
}

func (client *WSClient) messageReader() {
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client

			var newEvent = WSPaylaod{
				From: client.Username,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			return
		}
		var payload map[string]any
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}

		wsEvent := WSPaylaod{
			From: client.Username,
			Type: WS_MESSAGE_EVENT,
			Data: payload,
		}

		WSHub.HandleEvent(wsEvent)
	}
}

func (client *WSClient) messageWriter() {
	// ticker := time.NewTicker(pingPeriod)
	// defer func() {
	// 	ticker.Stop()
	// 	client.WSCoon.Close()
	// }()

	for {
		select {
		case message := <-client.OutgoingMsg:
			data, err := json.Marshal(message)
			if err != nil {
				return
			}
			err = client.WSCoon.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
			// case <-ticker.C:
			// 	client.WSCoon.SetWriteDeadline(time.Now().Add(writeWait))

			// 	if err := client.WSCoon.WriteMessage(websocket.PingMessage, nil); err != nil {
			// 		return
			// 	}
		}
	}
}
