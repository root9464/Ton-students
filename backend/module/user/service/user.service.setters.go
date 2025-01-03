package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/utils"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func (s *userService) Create(ctx context.Context, dto *user_dto.CreateUserDto) (*user_model.User, error) {
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
		Username:     initData.User.Username,
		Firstname:    initData.User.FirstName,
		Lastname:     initData.User.LastName,
		SelectedName: "username",
		Role:         "user",
		IsPremium:    initData.User.IsPremium,
		Hash:         initData.Hash,
	}

	modelUser, err := utils.ConvertDtoToEntity(&srcUser, user_model.User{})
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("creating user: %+v", modelUser)

	userInDb, err := s.repo.GetByID(ctx, modelUser.ID)
	if err != nil {
		s.logger.Warnf("get user by id error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if userInDb != nil {
		newUser, err := s.repo.Create(ctx, modelUser)
		if err != nil {
			s.logger.Warnf("create user error: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: err.Error(),
			}
		}

		s.logger.Infof("created user: %+v", newUser)

		return newUser, nil
	}

	s.logger.Infof("user already exists: %+v", modelUser)

	updateUser, err := s.repo.Update(ctx, modelUser)
	if err != nil {
		s.logger.Warnf("update user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return updateUser, nil
}
