package service

import (
	"forum/config"
	"forum/models"
	r "forum/server/repositories"
	"time"

	"github.com/gofrs/uuid/v5"
)

type ChatService struct {
	ChatRepo r.ChatRepository
}

func (chatService *ChatService) init() {
	chatService.ChatRepo = r.ChatRepo
}

func (chatService *ChatService) NewChat(chat models.Chat) error {
	chatId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	chat.ChatId = chatId
	chat.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err = chatService.ChatRepo.SaveChat(chat)
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

func (chatService *ChatService) GetChat(chatId string) (models.Chat, error) {
	chat, err := chatService.ChatRepo.GetChat(chatId)
	if err != nil {
		return models.Chat{}, err
	}
	return chat, nil
}

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

func (chatService *ChatService) AddUserToChat(userChat models.UserChat) error {
	err := chatService.ChatRepo.AddUserToChat(userChat.ChatId, userChat.UserId)
	if err != nil {
		return err
	}
	return nil
}
