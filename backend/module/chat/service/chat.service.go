package chat_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	chat_dto "github.com/root9464/Ton-students/module/chat/dto"
	chat_repository "github.com/root9464/Ton-students/module/chat/repository"
	serv_service "github.com/root9464/Ton-students/module/service_module/service"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IChatService = (*chatService)(nil)

type IChatService interface {
	CreateChat(ctx context.Context, dto *chat_dto.CreateChatType) (*string, error)
	GetChatIDBetweenUsers(ctx context.Context, userIDs []int64) (*string, error)
	CreateOrLoadChat(ctx context.Context, dto *chat_dto.CreateOrLoad) (*string, error)
}

type chatService struct {
	logger    *logger.Logger
	validator *validator.Validate

	serviceService serv_service.IServiceModuleService
	repo           chat_repository.IChatRepository
}

func NewChatService(
	logger *logger.Logger,
	vavalidator *validator.Validate,
	repo chat_repository.IChatRepository,
	serviceService serv_service.IServiceModuleService,
) *chatService {
	return &chatService{
		logger:         logger,
		validator:      vavalidator,
		repo:           repo,
		serviceService: serviceService,
	}
}
