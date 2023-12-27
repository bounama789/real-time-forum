package repositories

import (
	"forum/config"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
	"time"
)

type ChatRepository struct {
	BaseRepo
}

func (r *ChatRepository) init() {
	r.DB = db.DB
	r.TableName = db.CHATS_TABLE
}

func (r *ChatRepository) SaveChat(chat models.Chat) error {
	chat.CreatedAt = time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err := r.DB.Insert(r.TableName, chat)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) DeleteChat(chatId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"cht_id": chatId})
	if err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) GetChat(chatId string) (chat models.Chat, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"cht_id": chatId})
	if err != nil {
		return chat, err
	}
	err = row.Scan(&chat.ChatId, &chat.CreatedAt)
	if err != nil {
		return chat, err
	}
	return chat, nil
}

func (r *ChatRepository) GetAllChats() (chats []models.Chat, err error) {
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"1": 1}, "created_at DESC")
	if err != nil {
		return chats, err
	}
	for rows.Next() {
		var chat models.Chat
		err = rows.Scan(&chat.ChatId, &chat.CreatedAt)
		if err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
