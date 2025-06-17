package repositories

import (
	"backend/model"
	"backend/model/dto"
)

type FriendsRepository interface {
	FindFriend(friendsId uint) (*model.Friend, error)
	GetFriendUserByNumber(number int) (*model.User, error)
	SetRoomFriend() (uint, error)
	AddFriend(user *model.Friend, friend *model.Friend) error
	UpdateFriend(friendsId uint, friend *model.Friend) error
	DeleteFriend(friendsId uint) error
	GetAllFriends(userId uint) ([]*dto.FriendsUser, error)
}
