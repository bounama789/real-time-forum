package handler

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"forum/server/service"
	"io"
	"net/http"
	"strings"

	"github.com/gofrs/uuid/v5"
)

func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				fmt.Println(err)
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		// Unmarshal request body into comment struct
		var message models.Message
		err = json.Unmarshal(content, &message)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// Set SenderId from token data
		message.SenderId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Validate message body
		if strings.TrimSpace(message.Body) == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// create message in database
		err = service.MessService.NewMessage(message)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithJSON(w, http.StatusOK, message)
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func EditMessageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				fmt.Println(err)
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		// Unmarshal request body into comment struct
		var message models.Message
		err = json.Unmarshal(content, &message)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// Set SenderId from token data
		message.SenderId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Validate message body
		if strings.TrimSpace(message.Body) == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// create message in database
		err = service.MessService.EditMessage(message)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithJSON(w, http.StatusOK, message)
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		messageId := r.URL.Query().Get("messageId")
		if messageId == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// Read request body
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			if err == io.EOF {
				RespondWithError(w, http.StatusBadRequest, "Bad Request")
				return
			} else {
				fmt.Println(err)
				RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}

		// Unmarshal request body into comment struct
		var message models.Message
		err = json.Unmarshal(content, &message)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		// Set SenderId from token data
		message.SenderId, err = uuid.FromString(tokenData.UserId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Validate message body
		if strings.TrimSpace(message.Body) == "" {
			RespondWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// create message in database
		err = service.MessService.DeleteMessage(messageId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithJSON(w, http.StatusOK, message)
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Get token data from authentication cookie
		cookie, err := r.Cookie("auth-cookie")
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		tokenData, err := service.AuthSrvice.GetTokenData(cookie.Value)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Get messages from database
		messages, err := service.MessService.GetMessage(tokenData.UserId)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		} else {
			RespondWithJSON(w, http.StatusOK, messages)
		}
	default:
		RespondWithError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
