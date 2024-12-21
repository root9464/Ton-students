package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/root9464/Ton-students/ent"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	"github.com/root9464/Ton-students/shared/utils"
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

	srcUser := user_dto.SrcUser{
		FirstName:    initData.User.FirstName,
		ID:           initData.User.ID,
		IsBot:        initData.User.IsBot,
		IsPremium:    initData.User.IsPremium,
		LastName:     initData.User.LastName,
		UserName:     initData.User.Username,
		LanguageCode: initData.User.LanguageCode,
		PhotoURL:     initData.User.PhotoURL,
		Hash:         initData.Hash,
	}

	modelUser := new(ent.User)
	if err := utils.DtoToModel(&srcUser, modelUser); err != nil {
		return nil, err
	}

	s.logger.Infof("%v", modelUser)

	user, err := s.repo.Update(ctx, modelUser)
	if err != nil {
		if err.Error() == "user not found" {
			s.logger.Info("User not found create")
			user, err := s.repo.Create(ctx, modelUser)
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

// func (s *userService) UpdateInfo(ctx context.Context, id int64, dto *user_dto.UpdateUserDto) (*ent.User, error) {
// 	if err := s.validator.Struct(dto); err != nil {
// 		s.logger.Warnf("validate error: %s", err.Error())
// 		return nil, &fiber.Error{
// 			Code:    400,
// 			Message: err.Error(),
// 		}
// 	}
//
// 	user, err := s.repo.Update(ctx, id, dto)
// 	if err != nil {
// 		s.logger.Warnf("update user error: %s", err.Error())
// 		return nil, &fiber.Error{
// 			Code:    500,
// 			Message: err.Error(),
// 		}
// 	}
//
// 	return nil, nil
// }
