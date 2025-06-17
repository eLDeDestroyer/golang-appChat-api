package config

import (
	"backend/app/controllers"
	"backend/app/repositories"
	"backend/app/services"
	"gorm.io/gorm"
)

func DependencyInjection(db *gorm.DB) (*controllers.UserControllerImpl, *controllers.FriendControllerImpl, *controllers.ChatsControllerImpl) {
	userRepository := repositories.NewUserRepositoriesImpl(db)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserControllerImpl(userService)

	friendRepository := repositories.NewFriendRepositoriesImpl(db)
	friendService := services.NewFriendServiceImpl(friendRepository)
	friendController := controllers.NewFriendControllerImpl(friendService)

	chatRepository := repositories.NewChatRepositoriesImpl(db)
	chatService := services.NewChatsServiceImpl(chatRepository)
	chatController := controllers.NewChatsControllerImpl(chatService)

	return userController, friendController, chatController
}
