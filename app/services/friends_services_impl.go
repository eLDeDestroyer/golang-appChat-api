package services

import (
	"backend/app/repositories"
	"backend/model"
	"backend/model/dto"
)

type FriendServiceImpl struct {
	friendRepository repositories.FriendsRepository
}

func NewFriendServiceImpl(friendRepository repositories.FriendsRepository) *FriendServiceImpl {
	return &FriendServiceImpl{
		friendRepository: friendRepository,
	}
}

func (service *FriendServiceImpl) FindFriendUser(room_id uint, user_id uint) (*model.Friend, error) {
	return service.friendRepository.FindFriend(room_id, user_id)
}

func (service *FriendServiceImpl) AddFriendUser(number int, userId uint, name string) error {
	var friend model.Friend
	var user model.Friend

	roomId, err := service.friendRepository.SetRoomFriend()
	data, err := service.friendRepository.GetFriendUserByNumber(number)
	if err != nil {
		return err
	}

	friend.Friend = userId
	friend.UserId = data.Id
	friend.Name = ""
	friend.RoomId = roomId

	user.Friend = data.Id
	user.UserId = userId
	user.Name = name
	user.RoomId = roomId

	err = service.friendRepository.AddFriend(&user, &friend)
	if err != nil {
		return err
	}

	return nil
}

func (service *FriendServiceImpl) UpdateFriendUser(friendsId uint, friend *model.Friend) error {
	return service.friendRepository.UpdateFriend(friendsId, friend)
}

func (service *FriendServiceImpl) DeleteFriendUSer(friendsId uint) error {
	return service.friendRepository.DeleteFriend(friendsId)
}

func (service *FriendServiceImpl) GetAllFriendsUser(userId uint) ([]*dto.FriendsUser, error) {
	return service.friendRepository.GetAllFriends(userId)
}
