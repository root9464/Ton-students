package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/root9464/Ton-students/ent"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *userService) Create(ctx context.Context, dto interface{}) (*ent.User, error) {
	var data user_dto.CreateUserDto
	err := mapstructure.Decode(dto, &data)
	if err != nil {
		s.logger.Warnf("decode error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	initData, err := initdata.Parse(data.InitDataRaw)
	if err != nil {
		s.logger.Warnf("parse init data error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	s.logger.Infof("initData: %v", initData.User.ID)

	user, err := s.repo.Update(ctx, &initData)
	if err != nil {
		if err.Error() == "user not found" {
			s.logger.Info("User not found create")
			user, err := s.repo.Create(ctx, &initData)
			if err != nil {
				s.logger.Warnf("create user error: %s", err.Error())
				return nil, &fiber.Error{
					Code:    500,
					Message: err.Error(),
				}
			}
			return user, nil
		}
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return user, nil
}
