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
	CreateOrLoad(ctx *fiber.Ctx) error
}

type ChatController struct {
	logger      *logger.Logger
	chatService chat_service.IChatService

	connections map[int64]string // Мапа для хранения подключений (ключ — идентификатор пользователя)
	userToChat  map[string]string
	chatToUsers map[string][]string
	mu          sync.RWMutex
}

func NewChatController(logger *logger.Logger, chatService chat_service.IChatService) *ChatController {
	return &ChatController{
		logger:      logger,
		chatService: chatService,
		connections: make(map[int64]string),
		userToChat:  make(map[string]string),
		chatToUsers: make(map[string][]string),
		mu:          sync.RWMutex{},
	}
}
