package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *userService) Create(ctx context.Context, dto *user_dto.CreateUserDto) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	initData, err := initdata.Parse(dto.InitDataRaw)
	if err != nil {
		s.logger.Warnf("parse init data error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	user, err := s.repo.GetByID(ctx, initData.User.ID)
	if err != nil {
		return err
	}

	if user != nil {
		if err := s.repo.Update(ctx, &initData); err != nil {
			return err
		}
		return nil
	}

	return nil
}
