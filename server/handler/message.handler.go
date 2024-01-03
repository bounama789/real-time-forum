package handler

import (
	"forum/models"
	"forum/server/service"
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			if client.ChatId == msg.ChatId {
				err := client.WriteJSON(msg)
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer ws.Close()

	// Obtenez le ChatId associé à cette connexion (par exemple, à partir des paramètres de l'URL)
	chatId := r.URL.Query().Get("chatId")

	// Assurez-vous de récupérer également l'ID de l'utilisateur (par exemple, à partir du token)
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	client := models.UsersChats{
		Conn:   ws,
		UserId: tokenData.UserId,
		ChatId: uuid.FromStringOrNil(chatId),
	}

	err = service.ChatSrvice.AddUserToChat(client)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}
