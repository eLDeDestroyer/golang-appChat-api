package services

import (
	"backend/model"
	"backend/model/dto"
)

type UserService interface {
	GetUser(id uint) (*model.User, error)
	RegisterUser(user *model.User) error
	LoginUser(username string, password string) (*dto.ResponseWithToken, error)
	UpdateUser(id uint, user *model.User) error
}
