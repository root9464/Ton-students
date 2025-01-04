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

	s.logger.Infof("validate success: %+v", dto)

	if dto.SelectedName == nil {
		s.logger.Warnf("selected name is nil")
		selectedName := user_model.SelectedName("username")
		dto.SelectedName = &selectedName
	}

	modelUser, err := utils.ConvertDtoToEntity[user_model.User](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("creating user: %+v", modelUser)

	userInDb, err := s.repo.GetByID(ctx, dto.ID)
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
			Status:  "success",
			Message: "User created successfully",
			Data: &ResponseUserData{
				ID:           newUser.ID,
				Visiblename:  userVisibleName,
				SelectedName: newUser.SelectedName,
				Role:         newUser.Role,
				Infos:        newUser.Infos,
				IsPremium:    newUser.IsPremium,
				Hash:         newUser.Hash,
			},
		}, nil
	}

	if userInDb.Hash != modelUser.Hash {
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

		s.logger.Infof("updated user: %+v", updateUser)

		return &ResponseCreateUser{
			Status:  "success",
			Message: "User updated successfully",
			Data: &ResponseUserData{
				ID:           updateUser.ID,
				Visiblename:  userVisibleName,
				SelectedName: user_model.SelectedName(userVisibleName),
				Role:         updateUser.Role,
				Infos:        updateUser.Infos,
				IsPremium:    updateUser.IsPremium,
				Hash:         updateUser.Hash,
			},
		}, nil
	}

	userVisibleName := user_funcs.GetVisibleName(userInDb)

	return &ResponseCreateUser{
		Status:  "success",
		Message: "User get successfully",
		Data: &ResponseUserData{
			ID:           userInDb.ID,
			Visiblename:  userVisibleName,
			SelectedName: userInDb.SelectedName,
			Role:         userInDb.Role,
			Infos:        userInDb.Infos,
			IsPremium:    userInDb.IsPremium,
			Hash:         userInDb.Hash,
		},
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

func (s *userService) SelectVisibleName(ctx context.Context, dto *user_dto.SelectVisibleNameType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	userInDb, err := s.repo.GetByID(ctx, dto.ID)
	if err != nil {
		s.logger.Warnf("get user by id error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if userInDb == nil {
		return &fiber.Error{
			Code:    404,
			Message: "User not found",
		}
	}

	userInDb.SelectedName = dto.SelectedName
	userInDb.Hash = dto.Hash

	_, err = s.repo.Update(ctx, userInDb)
	if err != nil {
		s.logger.Warnf("update user error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
