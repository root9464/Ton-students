package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (s *userService) GetUser(ctx context.Context, id int64) (*user_dto.ShortUserType, error) {

	user, err := s.repo.GetByID(ctx, id)
	if err != nil || user == nil {
		return nil, &fiber.Error{
			Code:    404,
			Message: "User not found",
		}
	}

	convert, err := utils.ConvertDtoToEntity[user_dto.ShortUserType](user)
	if err != nil {
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return convert, nil
}
