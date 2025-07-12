package repositories

import (
	"backend/model"
	"backend/model/dto"
	"errors"
	"gorm.io/gorm"
)

type FriendRepositoriesImpl struct {
	db *gorm.DB
}

func NewFriendRepositoriesImpl(db *gorm.DB) *FriendRepositoriesImpl {
	return &FriendRepositoriesImpl{
		db: db,
	}
}

func (repo *FriendRepositoriesImpl) FindFriend(room_id uint, user_id uint) (*model.Friend, error) {
	var friend model.Friend

	err := repo.db.Table("friends").Where("room_id = ? AND user_id = ?", room_id, user_id).First(&friend).Error
	if err != nil {
		return nil, errors.New("friend not found")
	}

	return &friend, nil
}

func (repo *FriendRepositoriesImpl) GetFriendUserByNumber(number int) (*model.User, error) {
	var user model.User
	err := repo.db.Table("users").Where("unique_number = ?", number).First(&user).Error
	if err != nil {
		return &user, errors.New("friend not found")
	}

	return &user, nil
}

func (repo *FriendRepositoriesImpl) SetRoomFriend() (uint, error) {
	var room *model.Room
	err := repo.db.Table("rooms").Create(&room).Error
	if err != nil {
		return 0, err
	}

	return room.Id, nil
}

func (repo *FriendRepositoriesImpl) AddFriend(user *model.Friend, friend *model.Friend) error {
	var existsFriend model.Friend

	err := repo.db.Table("friends").
		Where("user_id = ?", friend.UserId).
		Where("friend = ?", friend.Friend).
		First(&existsFriend).Error

	if err == nil {
		return errors.New("friend already exists")
	}

	err = repo.db.Table("friends").Create(&user).Error
	err = repo.db.Table("friends").Create(&friend).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *FriendRepositoriesImpl) UpdateFriend(friendsId uint, friend *model.Friend) error {
	err := repo.db.Table("friends").
		Where("id = ?", friendsId).
		Updates(&friend).Error

	if err != nil {
		return errors.New("cannot update friend")
	}

	return nil
}

func (repo *FriendRepositoriesImpl) DeleteFriend(friendsId uint) error {
	var friend *model.Friend

	err := repo.db.Table("friends").Where("id = ?", friendsId).First(&friend).Error
	if err != nil {
		return errors.New("friend not found")
	}

	err = repo.db.Table("rooms").Where("id = ?", friend.RoomId).Delete(&model.Room{}).Error
	if err != nil {
		return errors.New("cannot delete room")
	}
	return nil
}

func (repo *FriendRepositoriesImpl) GetAllFriends(userId uint) ([]*dto.FriendsUser, error) {
	var friends []*dto.FriendsUser

	err := repo.db.Table("friends").
		Joins("JOIN users ON users.id = friends.friend").
		Select("users.unique_number as unique_number, friends.name as name, friends.room_id as room_id, users.id as id").
		Where("friends.user_id = ?", userId).
		Find(&friends).Error

	if err != nil {
		return nil, err
	}

	return friends, nil
}
