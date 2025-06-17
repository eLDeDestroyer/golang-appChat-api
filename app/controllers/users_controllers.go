package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}
