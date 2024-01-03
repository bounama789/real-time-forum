package service

import (
	"forum/config"
	"forum/models"
	r "forum/server/repositories"
	"time"

	"github.com/gofrs/uuid/v5"
)

type MessageService struct {
	MessageRepo r.MessageRepository
}

func (messageService *MessageService) init() {
	messageService.MessageRepo = r.MessRepo
}

func (messageService *MessageService) NewMessage(message models.Message) error {
	messageId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	message.MessId = messageId
	message.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err = messageService.MessageRepo.SaveMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (messageService *MessageService) DeleteMessage(messageId string) error {
	err := messageService.MessageRepo.DeleteMessage(messageId)
	if err != nil {
		return err
	}
	return nil
}

func (messageService *MessageService) EditMessage(updatedMessage models.Message) error {
	err := messageService.MessageRepo.UpdateMessage(updatedMessage)
	if err != nil {
		return err
	}
	return nil
}

func (messageService *MessageService) GetMessage(messageId string) (models.Message, error) {
	message, err := messageService.MessageRepo.GetMessage(messageId)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}

func (messageService *MessageService) GetChatMessages(ChatId string) ([]models.Message, error) {
	messages, err := messageService.MessageRepo.GetChatMessages(ChatId)
	if err != nil {
		return []models.Message{}, err
	}
	return messages, nil
}
