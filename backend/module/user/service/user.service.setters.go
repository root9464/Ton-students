package user_service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *userService) UpsertUser(ctx context.Context, dto *user_dto.UserType) (*user_dto.ShortUserType, error) {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	s.logger.Infof("validate success: %+v", dto)

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
		returnedUser, err := utils.ConvertDtoToEntity[user_dto.ShortUserType](newUser)
		returnedUser.VisibleName = utils.GetVisibleName(newUser)

		if err != nil {
			s.logger.Warnf("convert dto to entity error: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: err.Error(),
			}
		}

		return returnedUser, nil
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

		s.logger.Infof("updated user: %+v", updateUser)

		returnedUser, err := utils.ConvertDtoToEntity[user_dto.ShortUserType](updateUser)
		returnedUser.VisibleName = utils.GetVisibleName(updateUser)

		if err != nil {
			s.logger.Warnf("convert dto to entity error: %s", err.Error())
			return nil, &fiber.Error{
				Code:    500,
				Message: err.Error(),
			}
		}
		return returnedUser, nil
	}

	returnedUser, err := utils.ConvertDtoToEntity[user_dto.ShortUserType](userInDb)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	returnedUser.VisibleName = utils.GetVisibleName(userInDb)

	return returnedUser, nil
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

	userInfo, err := utils.ConvertDtoToEntity[user_model.UserInfo](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	s.logger.Infof("converted user info: %+v", userInfo)

	userInDb, err := s.repo.GetByID(ctx, userInfo.UserID)
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

	if len(*userInDb.Infos) >= 3 {
		return &fiber.Error{
			Code:    409,
			Message: "Maximum number of user infos reached",
		}
	}

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

	userInDb.SelectedName = user_model.SelectedName(dto.SelectedName)
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

func (s *userService) SetUserNickname(ctx context.Context, dto *user_dto.SetUserNicknameType) error {
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

	s.logger.Infof("set user nickname: %+v, user: %+v", dto.Nickname, *userInDb.Nickname)
	if *userInDb.Nickname == dto.Nickname {
		return &fiber.Error{
			Code:    400,
			Message: "Nickname don't change",
		}
	}

	userUpdate, err := utils.ConvertDtoToEntity[user_model.User](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	_, err = s.repo.Update(ctx, userUpdate)
	if err != nil {
		s.logger.Warnf("update user error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *userService) UpdateUserInfo(ctx context.Context, dto *user_dto.UpdateUserInfoType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	userInfo, err := utils.ConvertDtoToEntity[user_model.UserInfo](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if err := s.repo.UpdateUserInfo(ctx, userInfo); err != nil {
		s.logger.Warnf("update user info error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *userService) DeleteUserInfo(ctx context.Context, dto *user_dto.DeleteUserInfoType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := s.repo.DeleteUserInfo(ctx, dto.ID); err != nil {
		s.logger.Warnf("delete user info error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *userService) AddManyUserInfo(ctx context.Context, dto *user_dto.ManyUserInfoType) error {
	s.logger.Infof("Validating UserInfo slice: %+v", dto)
	if err := s.validator.RegisterValidation("maxinfos", user_dto.MaxInfos); err != nil {
		s.logger.Errorf("Error registering validator: %v", err)
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if err := s.validator.Struct(dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				s.logger.Errorf("Validation error: %s", e)
				return &fiber.Error{
					Code:    422,
					Message: fmt.Sprintf("validation error: %s exceeds maximum allowed entries (3)", e.Field()),
				}
			}
		}
		return err
	}

	s.logger.Infof("Get user by id: %+v", dto.UserId)

	userInDb, err := s.repo.GetByID(ctx, dto.UserId)
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

	if len(*userInDb.Infos) >= 3 {
		return &fiber.Error{
			Code:    409,
			Message: "Maximum number of user infos reached",
		}
	}

	recordsToAdd := 3 - len(*userInDb.Infos)
	newInfos := dto.Infos[:recordsToAdd]

	userManyInfos := make([]*user_model.UserInfo, len(newInfos))
	for i, info := range newInfos {
		userInfo, err := utils.ConvertDtoToEntity[user_model.UserInfo](&info)
		if err != nil {
			s.logger.Warnf("convert dto to entity error: %s", err.Error())
			return &fiber.Error{
				Code:    500,
				Message: err.Error(),
			}
		}

		userInfo.UserID = dto.UserId
		userManyInfos[i] = userInfo
	}

	for i, info := range userManyInfos {
		fmt.Printf("UserInfo[%d]: %+v\n", i, info)
	}

	if err := s.repo.AddManyUserInfo(ctx, userManyInfos); err != nil {
		s.logger.Warnf("add many user info error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}
