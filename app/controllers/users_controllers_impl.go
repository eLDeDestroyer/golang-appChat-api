package controllers

import (
	"backend/app/services"
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserControllerImpl(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (controller *UserControllerImpl) GetUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)

	user, err := controller.userService.GetUser(userId)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(ctx, user, "Success Get User")
}

func (controller *UserControllerImpl) RegisterUser(ctx *fiber.Ctx) error {
	var user model.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.userService.RegisterUser(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "User created successfully"})
}

func (controller *UserControllerImpl) LoginUser(ctx *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.BodyParser(&request)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	data, err := controller.userService.LoginUser(request.Username, request.Password)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(ctx, data, "Success Login")
}

func (controller *UserControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	var user model.User

	userId := ctx.Locals("userId").(uint)
	err := ctx.BodyParser(&user)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = controller.userService.UpdateUser(userId, &user)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(ctx, true, "Success UpdateUser")
}
