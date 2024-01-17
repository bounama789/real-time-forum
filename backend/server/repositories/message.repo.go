package repositories

import (
	db "forum/backend/database"
	opt "forum/backend/database/operators"
	q "forum/backend/database/query"
	"forum/backend/models"
)

type MessageRepository struct {
	BaseRepo
}

func (r *MessageRepository) init() {
	r.DB = db.DB
	r.TableName = db.MESSAGES_TABLE
}

func (r *MessageRepository) SaveMessage(message models.Message) error {
	err := r.DB.Insert(r.TableName, message)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessage(messageId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"mess_id": opt.Equals(messageId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) UpdateMessage(message models.Message) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"mess_id": opt.Equals(message.MessId)})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) GetMessage(messageId string) (message models.Message, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"mess_id": opt.Equals(messageId)})
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
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"cht_id": opt.Equals(ChatId)}, "created_at DESC")
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
