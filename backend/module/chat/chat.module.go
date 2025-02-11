package chat_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
	chat_controller "github.com/root9464/Ton-students/module/chat/controller"
	chat_repository "github.com/root9464/Ton-students/module/chat/repository"
	chat_service "github.com/root9464/Ton-students/module/chat/service"
	serv_service "github.com/root9464/Ton-students/module/service_module/service"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type ChatModule struct {
	chatService    chat_service.IChatService
	chatController chat_controller.IChatController
	chatRepo       chat_repository.IChatRepository

	serviceService serv_service.IServiceModuleService

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func NewChatModule(logger *logger.Logger, db *gorm.DB, vavalidator *validator.Validate, serviceService serv_service.IServiceModuleService) *ChatModule {
	return &ChatModule{
		logger:         logger,
		db:             db,
		validator:      vavalidator,
		serviceService: serviceService,
	}
}

func (m *ChatModule) ChatRepo() chat_repository.IChatRepository {
	if m.chatRepo == nil {
		m.chatRepo = chat_repository.NewChatRepository(m.db, m.logger)
	}
	return m.chatRepo
}

func (m *ChatModule) ChatService() chat_service.IChatService {
	if m.chatService == nil {
		m.chatService = chat_service.NewChatService(m.logger, m.validator, m.ChatRepo(), m.serviceService)
	}
	return m.chatService
}

func (m *ChatModule) ChatController() chat_controller.IChatController {
	if m.chatController == nil {
		m.chatController = chat_controller.NewChatController(m.logger, m.ChatService())
	}
	return m.chatController
}

func (m *ChatModule) ChatRoutes(router fiber.Router) {
	chat := router.Group("/chat")

	chat.Get("/conn/:id", socketio.New(m.ChatController().WS()))
	chat.Post("/", m.ChatController().CreateChat)
}
