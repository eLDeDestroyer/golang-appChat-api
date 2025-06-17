package services

import "backend/model"

type ChatServices interface {
	GetChatsUSer(id uint) ([]*model.Chat, error)
	AddChatUser(userId uint, chat *model.Chat) error
	DeleteAllChatsUser(id uint) error
}
