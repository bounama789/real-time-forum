package webs

import (
	"forum/server/handler"
	"forum/server/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Obtenez le ChatId associé à cette connexion (par exemple, à partir des paramètres de l'URL)
	chatId := r.URL.Query().Get("chatId")
	if chatId == "" {
		handler.RespondWithError(w, http.StatusBadRequest, "chatId is required")
		return
	}

	// Assurez-vous de récupérer également l'ID de l'utilisateur (par exemple, à partir du token)
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		handler.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	UserId := tokenData.UserId

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	hub := NewHub()
	go hub.Run()

	CreateNewSocketUser(hub, ws, UserId)
}

func GetAllChatsHandler(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous de récupérer également l'ID de l'utilisateur (par exemple, à partir du token)
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		handler.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	chats, err := service.ChatSrvice.GetAllChats(tokenData)
	if err != nil {
		handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, chats)
}
