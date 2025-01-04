package auth_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	auth_dto "github.com/root9464/Ton-students/module/auth/dto"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func (s *authService) Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*user_model.User, error) {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	initData, err := tma.Parse(dto.InitDataRaw)
	if err != nil {
		s.logger.Warnf("parse init data error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	srcUser := user_dto.UserType{
		ID:           initData.User.ID,
		Username:     initData.User.Username,
		Firstname:    &initData.User.FirstName,
		Lastname:     &initData.User.LastName,
		SelectedName: "username",
		Role:         "user",
		IsPremium:    initData.User.IsPremium,
		Hash:         initData.Hash,
	}

	userInDb, err := s.userService.Create(ctx, &srcUser)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return userInDb, nil
}
