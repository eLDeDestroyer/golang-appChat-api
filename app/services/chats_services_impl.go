package services

import (
	"backend/app/repositories"
	"backend/model"
)

type ChatsServiceImpl struct {
	chatRepository repositories.ChatsRepository
}

func NewChatsServiceImpl(chatRepository repositories.ChatsRepository) *ChatsServiceImpl {
	return &ChatsServiceImpl{
		chatRepository: chatRepository,
	}
}

func (service *ChatsServiceImpl) GetChatsUSer(id uint) ([]*model.Chat, error) {
	return service.chatRepository.GetChats(id)
}

func (service *ChatsServiceImpl) AddChatUser(userId uint, chat *model.Chat) error {
	chat.UserId = userId

	err := service.chatRepository.AddChat(chat)
	if err != nil {
		return err
	}
	return nil
}

func (service *ChatsServiceImpl) DeleteAllChatsUser(id uint) error {
	return service.chatRepository.DeleteAllChats(id)
}
