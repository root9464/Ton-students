package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *userService) Create(ctx context.Context, dto interface{}) error {
	var data user_dto.CreateUserDto
	err := mapstructure.Decode(dto, &data)
	if err != nil {
		s.logger.Error("Хуесос")
	}

	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	initData, err := initdata.Parse(data.InitDataRaw)
	if err != nil {
		s.logger.Warnf("parse init data error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	s.logger.Infof("initData: %v", initData.User.ID)

	if err := s.repo.Update(ctx, &initData); err != nil {
		if err.Error() == "user not found" {
			if err := s.repo.Create(ctx, &initData); err != nil {
				return &fiber.Error{
					Code:    500,
					Message: err.Error(),
				}
			}
		}
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
