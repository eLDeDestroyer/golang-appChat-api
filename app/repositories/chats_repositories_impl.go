package repositories

import (
	"backend/model"
	"gorm.io/gorm"
)

type ChatRepositoriesImpl struct {
	db *gorm.DB
}

func NewChatRepositoriesImpl(db *gorm.DB) *ChatRepositoriesImpl {
	return &ChatRepositoriesImpl{
		db: db,
	}
}

func (repo *ChatRepositoriesImpl) GetChats(id uint) ([]*model.Chat, error) {
	var chats []*model.Chat

	err := repo.db.Table("chats").Where("room_id = ?", id).Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (repo *ChatRepositoriesImpl) AddChat(chat *model.Chat) error {
	err := repo.db.Table("chats").Create(chat).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ChatRepositoriesImpl) DeleteAllChats(id uint) error {
	err := repo.db.Table("chats").Where("room_id = ?", id).Delete(&model.Chat{}).Error
	if err != nil {
		return err
	}
	return nil
}
