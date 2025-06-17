package helper

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, data any, message string) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
