package chat_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/shared/logger"
	"gorm.io/gorm"
)

type ChatModule struct {
	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
}

func NewChatModule(logger *logger.Logger, db *gorm.DB, vavalidator *validator.Validate) *ChatModule {
	return &ChatModule{
		logger:    logger,
		db:        db,
		validator: vavalidator,
	}
}

func (m *ChatModule) ChatRoutes(router fiber.Router) {

}
