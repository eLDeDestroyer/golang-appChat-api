package repositories

import (
	"backend/model"
	"gorm.io/gorm"
	"time"
)

type UserRepositoriesImpl struct {
	db *gorm.DB
}

func NewUserRepositoriesImpl(db *gorm.DB) *UserRepositoriesImpl {
	return &UserRepositoriesImpl{
		db: db,
	}
}

func (repo *UserRepositoriesImpl) GetMe(id uint) (*model.User, error) {
	var user model.User
	err := repo.db.Where("id = ?", id).
		Select("id", "unique_number", "name", "username").
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoriesImpl) Register(user *model.User) error {
	err := repo.db.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositoriesImpl) Login(username string) (*model.User, error) {
	var user model.User
	err := repo.db.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	timeNow := time.Now()
	user.LastLogin = &timeNow

	err = repo.db.Table("users").Where("username = ?", username).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoriesImpl) Update(id uint, user *model.User) error {
	err := repo.db.Table("users").Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}
