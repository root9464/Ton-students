package user_service

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/root9464/Ton-students/module/user/constant"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/utils"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *userService) Create(ctx context.Context, dto interface{}) error {
	var data user_dto.CreateUserDto
	err := mapstructure.Decode(dto, &data)
	if err != nil {
		s.logger.Warnf("decode error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
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

	srcUser := user_dto.SrcUser{
		FirstName:    initData.User.FirstName,
		ID:           initData.User.ID,
		IsBot:        initData.User.IsBot,
		IsPremium:    initData.User.IsPremium,
		UserName:     initData.User.Username,
		LanguageCode: initData.User.LanguageCode,
		PhotoURL:     initData.User.PhotoURL,
		Hash:         initData.Hash,
	}

	modelUser, err := utils.ConvertDtoToEntity[user_model.User](&srcUser)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("%v", modelUser)

	err = s.repo.Update(ctx, modelUser)
	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			s.logger.Info("User not found create")
			if err := s.repo.Create(ctx, modelUser); err != nil {
				s.logger.Warnf("create user error: %s", err.Error())
				return &fiber.Error{
					Code:    500,
					Message: err.Error(),
				}
			}
			return nil
		}
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
