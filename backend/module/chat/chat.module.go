package chat_module

import (
	"github.com/gofiber/fiber/v2"
	chat_controller "github.com/root9464/Ton-students/module/chat/controller"
	chat_service "github.com/root9464/Ton-students/module/chat/service"
)

type ChatModule struct {
	hub        *chat_service.Hub
	controller *chat_controller.ChatController
}

func NewChatModule() *ChatModule {
	hub := chat_service.NewHub()
	go hub.Run()
	controller := chat_controller.NewChatController(hub, chat_service.NewChatService(hub))
	return &ChatModule{hub: hub, controller: controller}
}

func (m *ChatModule) ChatRoutes(router fiber.Router) {
	router.Get("/chat", m.controller.HandleWebSocket)
}
