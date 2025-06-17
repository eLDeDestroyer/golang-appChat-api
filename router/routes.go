package router

import (
	"backend/app/controllers"
	"backend/app/middleware"
	"backend/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetUpRoutes(app *fiber.App, userController controllers.UserController, friendController controllers.FriendController, chatController controllers.ChatsController) {
	api := app.Group("/api")
	auth := app.Group("/api/auth", middleware.AuthMiddleware())
	api.Post("/users/register", userController.RegisterUser)
	api.Post("/users/login", userController.LoginUser)
	auth.Patch("/users/update", userController.UpdateUser)
	auth.Get("/users/me", userController.GetUser)

	auth.Get("/friend/:id", friendController.FindFriendUser)
	auth.Post("/friend/add", friendController.AddFriendUser)
	auth.Patch("/friend/update/:id", friendController.UpdateFriendUser)
	auth.Delete("/friend/delete", friendController.DeleteFriendUser)
	auth.Get("/friends", friendController.GetAllFriendsUser)

	auth.Get("/chats/:id", chatController.GetChatUser)
	auth.Post("/chats/add", chatController.AddChatUser)
	auth.Delete("/chats/delete/:id", chatController.DeleteAllChatUser)

	app.Get("/ws/chat", websocket.New(ws.WebSocketHub.HandleConnections))

}
