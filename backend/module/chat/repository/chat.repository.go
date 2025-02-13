package chat_repository

import (
	"context"

	chat_model "github.com/root9464/Ton-students/module/chat/model"
	common_model "github.com/root9464/Ton-students/module/model/common"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

var _ IChatRepository = (*chatRepository)(nil)

type IChatRepository interface {
	CreateChat(ctx context.Context, chat *chat_model.Chat) error
	CreateChatMembers(ctx context.Context, chatMembers []common_model.ChatUser) error
	GetChatIDBetweenUsers(ctx context.Context, userIDs []int64) (*string, error)
	GetChatByID(ctx context.Context, chatID string) (*chat_model.Chat, error)
}

type chatRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewChatRepository(db *gorm.DB, logger *logger.Logger) *chatRepository {
	return &chatRepository{
		db:     db,
		logger: logger,
	}
}
