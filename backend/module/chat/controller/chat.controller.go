package chat_controller

import (
	chat_service "github.com/root9464/Ton-students/module/chat/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type IChatController interface{}

type ChatController struct {
	logger      *logger.Logger
	chatService chat_service.IChatService
}

func NewChatController(logger *logger.Logger, chatService chat_service.IChatService) *ChatController {
	return &ChatController{
		logger:      logger,
		chatService: chatService,
	}
}
