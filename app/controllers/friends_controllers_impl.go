package controllers

import (
	"backend/app/services"
	"backend/helper"
	"backend/model"
	"github.com/gofiber/fiber/v2"
)

type FriendControllerImpl struct {
	friendService services.FriendsService
}

func NewFriendControllerImpl(friendService services.FriendsService) *FriendControllerImpl {
	return &FriendControllerImpl{
		friendService: friendService,
	}
}

func (controller *FriendControllerImpl) FindFriendUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("room_id")
	userId := ctx.Locals("userId").(uint)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	data, err := controller.friendService.FindFriendUser(uint(id), userId)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, data, "success get friend data")

}

func (controller *FriendControllerImpl) AddFriendUser(ctx *fiber.Ctx) error {
	var request struct {
		Name   string `json:"name"`
		Number int    `json:"unique_number"`
	}

	userId := ctx.Locals("userId").(uint)
	err := ctx.BodyParser(&request)

	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.friendService.AddFriendUser(request.Number, userId, request.Name)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, "success add friend user")
}

func (controller *FriendControllerImpl) UpdateFriendUser(ctx *fiber.Ctx) error {
	var friend model.Friend

	friendId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = ctx.BodyParser(&friend)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.friendService.UpdateFriendUser(uint(friendId), &friend)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, "success update friend user")
}

func (controller *FriendControllerImpl) DeleteFriendUser(ctx *fiber.Ctx) error {
	var request struct {
		Id uint `json:"id"`
	}

	err := ctx.BodyParser(&request)

	id := request.Id
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.friendService.DeleteFriendUSer(id)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, fiber.StatusOK, "success delete friend user")
}

func (controller *FriendControllerImpl) GetAllFriendsUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)

	data, err := controller.friendService.GetAllFriendsUser(userId)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(ctx, data, "success get friends user")
}
