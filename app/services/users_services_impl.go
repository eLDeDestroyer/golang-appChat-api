package services

import (
	"backend/app/repositories"
	"backend/helper"
	"backend/model"
	"backend/model/dto"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service *UserServiceImpl) GetUser(id uint) (*model.User, error) {
	return service.userRepository.GetMe(id)
}

func (service *UserServiceImpl) RegisterUser(user *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	rand.Seed(time.Now().UnixNano())	
	randomNumber := rand.Intn(90000000) + 10000000

	if err != nil {
		return err
	}

	user.UniqueNumber = randomNumber
	user.Password = string(hashPassword)
	return service.userRepository.Register(user)
}

func (service *UserServiceImpl) LoginUser(username string, password string) (*dto.ResponseWithToken, error) {
	var response dto.ResponseWithToken
	user, err := service.userRepository.Login(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	tokenString, err := helper.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	response.Token = tokenString
	response.Username = user.Username

	return &response, nil
}

func (service *UserServiceImpl) UpdateUser(id uint, user *model.User) error {
	return service.userRepository.Update(id, user)
}
