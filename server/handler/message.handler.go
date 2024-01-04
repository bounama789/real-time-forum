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

// broadcast can be used to send and receive values of type `models.Message`.
var broadcast = make(chan models.Message)

// Map clients pour stocker tous les clients connectés
var clients = make(map[*Client]bool)

// Structure Client pour représenter un client connecté
type Client struct {
	Conn   *websocket.Conn
	ChatId string // ou uuid.UUID selon la représentation utilisée dans votre système
	UserId uuid.UUID
}

// Fonction HandleConnections pour gérer les nouvelles connexions WebSocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Mise à niveau de la connexion en WebSocket
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

	// Créez une instance de UsersChats pour lier l'utilisateur au chat
	userChat := models.UsersChats{
		UserId: uuid.FromStringOrNil(tokenData.UserId),
		ChatId: uuid.FromStringOrNil(chatId),
	}

	// Enregistrez l'association utilisateur-chat dans votre service approprié
	err = service.ChatSrvice.AddUserToChat(userChat)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Créez une instance de la structure Client avec le ChatId
	client := &Client{
		Conn:   ws,
		ChatId: chatId,
		UserId: uuid.FromStringOrNil(tokenData.UserId),
	}

	// Ajoutez le client à la carte clients
	clients[client] = true

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			// Find the client associated with the websocket connection
			var client *Client
			for c := range clients {
				if c.Conn == ws {
					client = c
					break
				}
			}
			if client != nil {
				delete(clients, client)
			}
			break
		}
		broadcast <- msg
	}
}

// Fonction HandleMessages pour gérer l'envoi de messages à tous les clients
func HandleMessages() {
	for {
		msg := <-broadcast
		// Enregistrez le message dans votre service approprié
		SaveMessage(msg)
		for client := range clients {
			// Vérifiez si le client est dans le même chat que le message
			if uuid.FromStringOrNil(client.ChatId) == msg.ChatId {
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

// function for save message to database

func SaveMessage(msg models.Message) {
	err := service.MessService.NewMessage(msg)
	if err != nil {
		return
	}
}
