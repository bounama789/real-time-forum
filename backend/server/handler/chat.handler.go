package handler

import (
	"encoding/json"
	"forum/backend/server/cors"
	repo "forum/backend/server/repositories"
	"forum/backend/ws"
	"net/http"
)

func GetChats(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
		return
	}

	chats, err := repo.ChatRepo.GetUserChats(tokenData.UserId)
	if err != nil {
		println(err)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"msg": "no chat"})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(chats)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	tokenData, err := authService.VerifyToken(r)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"msg": "unauthorized"})
		return
	}

	users, err := repo.UserRepo.GetAllUsers()

	if err != nil {
		println(err)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"msg": "no users"})
		return
	}

	type reformatedUserData struct {
		Username string `json:"username"`
		Status   string `json:"status"`
	}

	var data []reformatedUserData

	for _, user := range users {
		if tokenData.Username == user.Username {
			continue
		}
		var status = "offline"

		if _,ok := ws.WSHub.Clients.Load(user.Username);ok{
			status = "online"
		}
		data = append(data, reformatedUserData{Username: user.Username, Status: status})
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)

}
