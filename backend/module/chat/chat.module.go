package chat_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
	chat_controller"github.com/root9464/Ton-students/module/chat/controller"
	"github.com/root9464/Ton-students/shared/logger"
)

type ChatModule struct {
	logger     *logger.Logger
	controller *chat_controller.ChatController
}

func NewChatModule(logger *logger.Logger) *ChatModule {
	chatController := chat_controller.NewChatController(logger)
	return &ChatModule{
		logger:     logger,
		controller: chatController,
	}
}

func (m *ChatModule) ChatRoutes(router fiber.Router) {
	router.Get("/createroom", m.controller.CreateRoomHandler)
	router.Get("/conn", websocket.New(m.controller.HandleWebSocket))
}
