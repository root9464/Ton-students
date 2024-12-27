package chat_module

import (
	"github.com/aws/aws-sdk-go/service"
	"github.com/gofiber/fiber/v2"
	//chat_controller
	//chat_service
)

type ChatModel struct {
	hub *chat_service.Hub
	controller *chat_controller.ChatController
}

func NewChatModule() *ChatModel {
	return nil
}