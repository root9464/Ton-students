package chat_service

import (
	"github.com/go-playground/validator/v10"
	chat_repository "github.com/root9464/Ton-students/module/chat/repository"
	"github.com/root9464/Ton-students/shared/logger"
)

var _ IChatService = (*chatService)(nil)

type IChatService interface{}

type chatService struct {
	logger    *logger.Logger
	validator *validator.Validate

	repo chat_repository.IChatRepository
}

func NewChatService(logger *logger.Logger, vavalidator *validator.Validate, repo chat_repository.IChatRepository) *chatService {
	return &chatService{
		logger:    logger,
		validator: vavalidator,
		repo:      repo,
	}
}
