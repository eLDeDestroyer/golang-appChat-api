package controllers

import "github.com/gofiber/fiber/v2"

type FriendController interface {
	FindFriendUser(ctx *fiber.Ctx) error
	AddFriendUser(ctx *fiber.Ctx) error
	UpdateFriendUser(ctx *fiber.Ctx) error
	DeleteFriendUser(ctx *fiber.Ctx) error
	GetAllFriendsUser(ctx *fiber.Ctx) error
}
