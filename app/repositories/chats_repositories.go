package repositories

import "backend/model"

type ChatsRepository interface {
	GetChats(id uint) ([]*model.Chat, error)
	AddChat(chat *model.Chat) error
	DeleteAllChats(id uint) error
}
