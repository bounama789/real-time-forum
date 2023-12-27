package service

import (
	"forum/models"
	r "forum/server/repositories"
)

type ChatService struct {
	ChatRepo r.ChatRepository
}

func (chatService *ChatService) init() {
	chatService.ChatRepo = r.ChatRepo
}

func (chatService *ChatService) NewChat(chat models.Chat) error {
	err := chatService.ChatRepo.SaveChat(chat)
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

func (chatService *ChatService) GetChats() ([]models.Chat, error) {
	chats, err := chatService.ChatRepo.GetAllChats()
	if err != nil {
		return []models.Chat{}, err
	}
	return chats, nil
}
