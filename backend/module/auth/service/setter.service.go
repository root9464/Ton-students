package auth_service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *authService) Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	botToken := s.config.TelegramBotToken
	expIn := 240 * time.Hour

	if err := initdata.Validate(dto.InitDataRaw, botToken, expIn); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := s.userService.Create(ctx, &user_dto.CreateUserDto{InitDataRaw: dto.InitDataRaw}); err != nil {
	}

	return nil
}