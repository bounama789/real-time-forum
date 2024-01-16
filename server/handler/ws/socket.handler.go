package webs

import (
	"bytes"
	"encoding/json"
	"forum/models"
	_ "forum/server/handler"
	"forum/server/service"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/websocket"
)

// The `const` block is defining constants for various time durations and message sizes used in the WebSocket implementation.
// The line `writeWait = 10 * time.Second` is defining a constant named `writeWait` with a value of 10 seconds. This constant is used to set the maximum time allowed for a write operation on a WebSocket
// connection. If a write operation takes longer than this duration, the connection will be closed. `pongWait = 60 * time.Second` is setting the maximum time allowed for a pong response from the client.
// The line `pingPeriod = (pongWait * 9) / 10` is calculating the duration between sending ping messages to the client.
// The line `maxMessageSize = 512` is setting the maximum size of a message that can be received or sent over the WebSocket connection. It ensures that the size of the message does not exceed 512
// bytes. If a message exceeds this size, it will be truncated or rejected.
const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// The function unregisters a client from a hub and closes its WebSocket connection.
func unRegisterAndCloseConnection(c *Client) {
	c.hub.unregister <- c
	err := c.webSocketConnection.Close()
	if err != nil {
		return
	}
}

// The function sets the read limit, read deadline, and pong handler for a WebSocket connection.
// setSocketPayloadReadConfig configures the WebSocket connection for reading messages. It sets the read limit, read deadline, and pong handler for the given client.
// The read limit is set to maxMessageSize, the read deadline is set to pongWait duration, and the pong handler is set to update the read deadline whenever a pong message is received.
func setSocketPayloadReadConfig(c *Client) {
	c.webSocketConnection.SetReadLimit(maxMessageSize)
	err := c.webSocketConnection.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	c.webSocketConnection.SetPongHandler(func(string) error {
		err := c.webSocketConnection.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})
}

// The function reads a message from a WebSocket connection.
// readSocketPayload reads a single JSON encoded message from the WebSocket connection and returns the message as a byte array.
func readSocketPayload(c *Client) ([]byte, error) {
	_, payload, err := c.webSocketConnection.ReadMessage()
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// The function sends a message to a WebSocket connection.
// writeSocketPayload writes a single JSON encoded message to the WebSocket connection.
func writeSocketPayload(c *Client, payload []byte) error {
	err := c.webSocketConnection.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	return c.webSocketConnection.WriteMessage(websocket.TextMessage, payload)
}

// The function handles the events sent by a client over the WebSocket
// connection. handleSocketPayloadEvents handles the events sent by a client over
// the WebSocket connection. It decodes the message payload and handles the event
// based on the event name. The message payload is decoded into a
// SocketEventStruct. The event name is extracted from the SocketEventStruct and
// the event is handled based on the event name. The events are handled by
// calling the appropriate function for the event. The events are handled by
// calling the appropriate function for the event. The `join` event is handled by
// the handleJoinEvent function, the `message` event is handled by the
// handleMessageEvent function, and the `disconnect` event is handled by the
// handleDisconnectEvent function. The handleSocketPayloadEvents function is
// called in a loop in the readPump function of the Client type. It is executed
// in a separate goroutine for each WebSocket connection. The function exits when
// the WebSocket connection is closed or an error occurs.

func handleSocketPayloadEvents(client *Client, socketEventPayload SocketEventStruct) {
	var socketEventResponse SocketEventStruct
	switch socketEventPayload.EventName {
	case "join":
		log.Printf("Join Event triggered")
		BroadcastSocketEventToAllClient(client.hub, SocketEventStruct{
			EventName: socketEventPayload.EventName,
			EventPayload: JoinDisconnectPayload{
				UserID: client.userID,
				Users:  getAllConnectedUsers(client.hub),
			},
		})

	case "disconnect":
		log.Printf("Disconnect Event triggered")
		BroadcastSocketEventToAllClient(client.hub, SocketEventStruct{
			EventName: socketEventPayload.EventName,
			EventPayload: JoinDisconnectPayload{
				UserID: client.userID,
				Users:  getAllConnectedUsers(client.hub),
			},
		})

	case "message":
		log.Printf("Message Event triggered")
		selectedUserID := socketEventPayload.EventPayload.(map[string]interface{})["userID"].(string)
		socketEventResponse.EventName = "message response"
		socketEventResponse.EventPayload = map[string]interface{}{
			"username": getUsernameByUserID(client.hub, selectedUserID),
			"message":  socketEventPayload.EventPayload.(map[string]interface{})["message"],
			"userID":   selectedUserID,
		}
		chatMessage := models.Message{
			ChatId:   uuid.FromStringOrNil(socketEventPayload.EventPayload.(map[string]interface{})["chat_id"].(string)),
			SenderId: uuid.FromStringOrNil(socketEventPayload.EventPayload.(map[string]interface{})["sender_id"].(string)),
			Body:     socketEventPayload.EventPayload.(map[string]interface{})["message"].(string),
		}
		if err := service.MessService.NewMessage(chatMessage); err != nil {
			log.Println(err)
		}
		EmitToSpecificClient(client.hub, socketEventResponse, selectedUserID)
	}
}

func getUsernameByUserID(hub *Hub, userID string) string {
	var username string
	for client := range hub.clients {
		if client.userID == userID {
			username = client.username
		}
	}
	return username
}

func getAllConnectedUsers(hub *Hub) []UserStruct {
	var users []UserStruct
	for singleClient := range hub.clients {
		users = append(users, UserStruct{
			Username: singleClient.username,
			UserID:   singleClient.userID,
		})
	}
	return users
}

func (c *Client) readPump() {
	var socketEventPayload SocketEventStruct

	defer unRegisterAndCloseConnection(c)

	setSocketPayloadReadConfig(c)

	for {
		_, payload, err := c.webSocketConnection.ReadMessage()

		// The code `decoder := json.NewDecoder(bytes.NewReader(payload))` creates a new JSON decoder that reads from a `bytes.Reader` containing the `payload` data.
		decoder := json.NewDecoder(bytes.NewReader(payload))
		decoderErr := decoder.Decode(&socketEventPayload)

		if decoderErr != nil {
			log.Printf("error: %v", decoderErr)
			break
		}

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error ===: %v", err)
			}
			break
		}

		handleSocketPayloadEvents(c, socketEventPayload)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.webSocketConnection.Close()
		if err != nil {
			return
		}
	}()
	for {
		select {
		// The line `case payload, ok := <-c.send:` is a case statement in a select block. It checks if there
		// is a value available to be received from the `c.send` channel.
		case payload, ok := <-c.send:
			// The code `reqBodyBytes := new(bytes.Buffer)` creates a new buffer to store the encoded JSON payload.
			reqBodyBytes := new(bytes.Buffer)
			if err := json.NewEncoder(reqBodyBytes).Encode(payload); err != nil {
				log.Printf("Failed to encode payload: %v", err)
				return
			}
			finalPayload := reqBodyBytes.Bytes()

			if err := c.webSocketConnection.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Failed to set write deadline: %v", err)
				return
			}
			if !ok {
				if err := c.webSocketConnection.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("Failed to write close message: %v\n", err)
					return
				}
				log.Printf("Failed to write close message: %v\n", ok)
				return
			}

			w, err := c.webSocketConnection.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Failed to get next writer: %v", err)
				return
			}

			if _, err = w.Write(finalPayload); err != nil {
				log.Printf("Failed to write payload: %v", err)
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				//  Encode the data from c.send into JSON and write it to reqBodyBytes.
				if err := json.NewEncoder(reqBodyBytes).Encode(<-c.send); err != nil {
					log.Printf("Failed to encode payload: %v", err)
					return
				}
				// writes the encoded JSON payload to the WebSocket connection.
				if _, err = w.Write(reqBodyBytes.Bytes()); err != nil {
					log.Printf("Failed to write payload: %v", err)
					return
				}
			}

			if err := w.Close(); err != nil {
				log.Printf("Failed to close writer: %v", err)
				return
			}
		case <-ticker.C:
			if err := c.webSocketConnection.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Failed to set write deadline: %v", err)
				return
			}
			if err := c.webSocketConnection.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Failed to write ping message: %v", err)
				return
			}
		}
	}
}

func CreateNewSocketUser(hub *Hub, connection *websocket.Conn, username string) {
	uniqueID := uuid.FromStringOrNil(username)
	client := &Client{
		hub:                 hub,
		webSocketConnection: connection,
		send:                make(chan SocketEventStruct),
		username:            username,
		userID:              uniqueID.String(),
	}

	go client.writePump()
	go client.readPump()

	client.hub.register <- client
}

// HandleUserRegisterEvent will handle the Join event for New socket users
func HandleUserRegisterEvent(hub *Hub, client *Client) {
	hub.clients[client] = true
	handleSocketPayloadEvents(client, SocketEventStruct{
		EventName:    "join",
		EventPayload: client.userID,
	})
}

// HandleUserDisconnectEvent will handle the Disconnect event for socket users
func HandleUserDisconnectEvent(hub *Hub, client *Client) {
	_, ok := hub.clients[client]
	if ok {
		delete(hub.clients, client)
		close(client.send)

		handleSocketPayloadEvents(client, SocketEventStruct{
			EventName:    "disconnect",
			EventPayload: client.userID,
		})
	}
}

// EmitToSpecificClient will emit the socket event to specific socket user
func EmitToSpecificClient(hub *Hub, payload SocketEventStruct, userID string) {
	for client := range hub.clients {
		if client.userID == userID {
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}

// BroadcastSocketEventToAllClient will emit the socket events to all socket users
func BroadcastSocketEventToAllClient(hub *Hub, payload SocketEventStruct) {
	for client := range hub.clients {
		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, client)
		}
	}
}
