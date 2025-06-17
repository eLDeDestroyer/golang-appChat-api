package controllers

import (
	"backend/app/services"
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

type ChatsControllerImpl struct {
	chatService services.ChatServices
}

func NewChatsControllerImpl(chatService services.ChatServices) *ChatsControllerImpl {
	return &ChatsControllerImpl{
		chatService: chatService,
	}
}

func (controller *ChatsControllerImpl) GetChatUser(ctx *fiber.Ctx) error {
	roomId, err := ctx.ParamsInt("id")

	data, err := controller.chatService.GetChatsUSer(uint(roomId))
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(ctx, data, "success")
}

func (controller *ChatsControllerImpl) AddChatUser(ctx *fiber.Ctx) error {
	var chat *model.Chat

	userId := ctx.Locals("userId").(uint)

	err := ctx.BodyParser(&chat)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}
	err = controller.chatService.AddChatUser(userId, chat)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, chat, "success")
}

func (controller *ChatsControllerImpl) DeleteAllChatUser(ctx *fiber.Ctx) error {
	roomId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.chatService.DeleteAllChatsUser(uint(roomId))
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, nil, "success")

}
