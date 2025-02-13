package chat_controller

import (
	"github.com/gofiber/fiber/v2"
	chat_dto "github.com/root9464/Ton-students/module/chat/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *ChatController) CreateOrLoad(ctx *fiber.Ctx) error {
	dto := new(chat_dto.CreateOrLoad)
	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON(&fiber.Error{
			Code:    400,
			Message: err.Error(),
		})
	}

	chatID, err := c.chatService.CreateOrLoadChat(ctx.Context(), dto)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Chat created or loaded successfully",
		"data":    chatID,
	})
}
