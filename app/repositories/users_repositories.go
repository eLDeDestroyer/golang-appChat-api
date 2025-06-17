package repositories

import "backend/model"

type UserRepository interface {
	GetMe(id uint) (*model.User, error)
	Register(user *model.User) error
	Login(username string) (*model.User, error)
	Update(id uint, user *model.User) error
}
