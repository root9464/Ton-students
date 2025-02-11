package chat_controller

import (
	"sync"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
	chat_service "github.com/root9464/Ton-students/module/chat/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type IChatController interface {
	WS() func(*socketio.Websocket)
	CreateChat(ctx *fiber.Ctx) error
}

type ChatController struct {
	logger      *logger.Logger
	chatService chat_service.IChatService

	connections map[int64]string   // Мапа для хранения подключений (ключ — идентификатор пользователя)
	chats       map[string][]int64 // Мапа для хранения чатов и их участников
	mu          sync.RWMutex
}

func NewChatController(logger *logger.Logger, chatService chat_service.IChatService) *ChatController {
	return &ChatController{
		logger:      logger,
		chatService: chatService,
		connections: make(map[int64]string),
		mu:          sync.RWMutex{},
	}
}
