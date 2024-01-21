package service

import (
	"forum/backend/config"
	"forum/backend/models"
	"forum/backend/server/repositories"
	r "forum/backend/server/repositories"
	"forum/backend/ws"
	"slices"
	"time"

	"github.com/gofrs/uuid/v5"
)

type ChatService struct {
	ChatRepo r.ChatRepository
}

func (chatService *ChatService) init() {
	chatService.ChatRepo = r.ChatRepo
}

func (chatService *ChatService) NewChat(chat *models.Chat) error {
	chatId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	chat.ChatId = chatId
	chat.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err = chatService.ChatRepo.SaveChat(*chat)
	if err != nil {
		return err
	}
	return nil
}

func (chatService *ChatService) UpdateChat(chat models.Chat) error {
	err := chatService.ChatRepo.UpdateChat(chat)
	if err != nil {
		return err
	}
	return nil
}

func (chatService *ChatService) DeleteChat(chatId string) error {
	err := chatService.ChatRepo.DeleteChat(chatId)
	if err != nil {
		return err
	}
	return nil
}

// func (chatService *ChatService) GetChat(chatId string) (models.Chat, error) {
// 	chat, err := chatService.ChatRepo.GetChat(chatId)
// 	if err != nil {
// 		return models.Chat{}, err
// 	}
// 	return chat, nil
// }

func (chatService *ChatService) GetAllChats(t models.TokenData) ([]models.Chat, error) {
	chats, err := chatService.ChatRepo.GetAllChats(t)
	if err != nil {
		return chats, err
	}
	return chats, nil
}

func (chatService *ChatService) GetChatMessages(chatId string) ([]models.Message, error) {
	messages, err := chatService.ChatRepo.GetChatMessages(chatId)
	if err != nil {
		return messages, err
	}
	return messages, nil
}

func (chatService *ChatService) GetChatStatus(username string) (any, error) {
	chats, err := repositories.ChatRepo.GetUserChats(username)
	if err != nil {
		println(err)
		return nil, err
	}

	type reformatedUserData struct {
		Username string `json:"username"`
		Status   string `json:"status"`
	}

	var data []reformatedUserData

	for _, chat := range chats {
		var uname string
		if username == chat.Recipient {
			uname = chat.Requester
		} else {
			uname = chat.Recipient
		}
		var status = "offline"
 		_, ok := ws.WSHub.Clients.Load(uname);
		if ok {
			status = "online"
		}
		data = append(data, reformatedUserData{Username: uname, Status: status})
	}

	users, err := repositories.UserRepo.GetAllUsers()

	if err != nil {

		return nil, err
	}

	for _, user := range users {
		if username == user.Username {
			continue
		}
		if slices.ContainsFunc(data, func(rud reformatedUserData) bool {
			return rud.Username == user.Username
		}) {
			continue
		}
		var status = "offline"

		if _, ok := ws.WSHub.Clients.Load(user.Username); ok {
			status = "online"
		}

		data = append(data, reformatedUserData{Username: user.Username, Status: status})
	}

	return data, nil
}

// func (chatService *ChatService) AddUserToChat(userChat models.) error {
// 	err := chatService.ChatRepo.AddUserToChat(userChat.ChatId, userChat.UserId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
