package service

import (
	"forum/models"
	r "forum/server/repositories"
)

type MessageService struct {
	MessageRepo r.MessageRepository
}

func (messageService *MessageService) init() {
	messageService.MessageRepo = r.MessRepo
}

func (messageService *MessageService) NewMessage(message models.Message) error {
	err := messageService.MessageRepo.SaveMessage(message)
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

