package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_funcs "github.com/root9464/Ton-students/module/user/funcs"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *userService) Create(ctx context.Context, dto *user_dto.UserType) (*ResponseCreateUser, error) {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	modelUser, err := utils.ConvertDtoToEntity(dto, user_model.User{})
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

	if userInDb == nil {
		newUser, err := s.repo.Create(ctx, modelUser)
		if err != nil {
			s.logger.Warnf("create user error: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: err.Error(),
			}
		}

		s.logger.Infof("created user: %+v", newUser)

		userVisibleName := user_funcs.GetVisibleName(newUser)

		return &ResponseCreateUser{
			ID:           newUser.ID,
			Visiblename:  userVisibleName,
			SelectedName: newUser.SelectedName,
			Role:         newUser.Role,
			IsPremium:    newUser.IsPremium,
			Hash:         newUser.Hash,
		}, nil
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

	userVisibleName := user_funcs.GetVisibleName(updateUser)

	return &ResponseCreateUser{
		ID:           updateUser.ID,
		Visiblename:  userVisibleName,
		SelectedName: updateUser.SelectedName,
		Role:         updateUser.Role,
		IsPremium:    updateUser.IsPremium,
		Hash:         updateUser.Hash,
	}, nil
}

func (s *userService) AddUserInfo(ctx context.Context, dto *user_dto.UserInfoType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	s.logger.Infof("adding user info: %+v", dto)

	userInfo := &user_model.UserInfo{
		UserID:  dto.UserId,
		Title:   dto.Title,
		Content: dto.Content,
	}

	s.logger.Infof("converted user info: %+v", userInfo)

	if err := s.repo.AddUserInfo(ctx, userInfo); err != nil {
		s.logger.Warnf("add user info error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
