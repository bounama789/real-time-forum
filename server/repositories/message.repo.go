package repositories

import (
	"forum/config"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
	"time"
)

type MessageRepository struct {
	BaseRepo
}

func (r *MessageRepository) init() {
	r.DB = db.DB
	r.TableName = db.MESSAGES_TABLE
}

func (r *MessageRepository) SaveMessage(message models.Message) error {
	message.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err := r.DB.Insert(r.TableName, message)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(messageId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"mess_id": messageId})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) UpdateMessage(message models.Message) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"mess_id": message.MessId})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetMessage(messageId string) (message models.Message, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"mess_id": messageId})
	if err != nil {
		return message, err
	}
	err = row.Scan(&message.MessId, &message.ChatId, &message.SenderId, &message.Body, &message.CreatedAt)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (r *MessageRepository) GetChatMessages(ChatId string) (messages []models.Message, err error) {
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"cht_id": ChatId}, "created_at DESC")
	if err != nil {
		return messages, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.MessId, &message.ChatId, &message.SenderId, &message.Body, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
