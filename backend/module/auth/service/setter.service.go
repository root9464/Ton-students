package auth_service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/ent"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *authService) Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*ent.User, error) {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	botToken := s.config.TelegramBotToken
	expIn := 24 * time.Hour

	if err := initdata.Validate(dto.InitDataRaw, botToken, expIn); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	user, err := s.userService.Create(ctx, dto)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return nil, err
	}

	return user, nil
}
