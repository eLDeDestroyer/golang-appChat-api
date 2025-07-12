package services

import (
	"backend/model"
	"backend/model/dto"
)

type FriendsService interface {
	FindFriendUser(room_id uint, user_id uint) (*model.Friend, error)
	AddFriendUser(number int, userId uint, name string) error
	UpdateFriendUser(friendsId uint, friend *model.Friend) error
	DeleteFriendUSer(friendsId uint) error
	GetAllFriendsUser(userId uint) ([]*dto.FriendsUser, error)
}
