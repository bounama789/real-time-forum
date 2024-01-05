package handler

import (
	"forum/models"
	"forum/server/service"
	"net/http"
)

func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	_, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		var chat models.Chat
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		err = service.ChatSrvice.NewChat(chat)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		RespondWithJSON(w, http.StatusOK, chat)
	}
}

func DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	_, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	switch r.Method {
	case http.MethodDelete:
		chatId := r.URL.Query().Get("chatId")
		if chatId == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		err := service.ChatSrvice.DeleteChat(chatId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		RespondWithJSON(w, http.StatusOK, "Chat deleted")
	}
}

func GetAllChatsHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	chats, err := service.ChatSrvice.GetAllChats(tokenData)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	RespondWithJSON(w, http.StatusOK, chats)
}

func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	_, err := service.AuthSrvice.VerifyToken(r)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	switch r.Method {
	case http.MethodGet:
		chatId := r.URL.Query().Get("chatId")
		if chatId == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		chat, err := service.ChatSrvice.GetChat(chatId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		RespondWithJSON(w, http.StatusOK, chat)
	}
}
