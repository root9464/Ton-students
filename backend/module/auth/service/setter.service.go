package auth_service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/module/auth/dto"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func (s *authService) Authorize(ctx context.Context, dto *dto.AutorizeDto) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	botToken := s.config.TelegramBotToken
	expIn := 240 * time.Hour

	if err := tma.Validate(dto.InitDataRaw, botToken, expIn); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	return nil
}
