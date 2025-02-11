package chat_controller

import (
	"github.com/gofiber/fiber/v2"
	chat_dto "github.com/root9464/Ton-students/module/chat/dto"
)

func (c *ChatController) CreateChat(ctx *fiber.Ctx) error {
	dto := new(chat_dto.CreateChatType)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := c.chatService.CreateChat(ctx.Context(), dto); err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"message": "Chat created",
	})
}
