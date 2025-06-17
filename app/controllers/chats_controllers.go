package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type ChatsController interface {
	GetChatUser(ctx *fiber.Ctx) error
	AddChatUser(ctx *fiber.Ctx) error
	DeleteAllChatUser(ctx *fiber.Ctx) error
}
