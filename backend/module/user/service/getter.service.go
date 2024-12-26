package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_model "github.com/root9464/Ton-students/module/user/model"
)

func (s *userService) GetByID(ctx context.Context, id int64) (*user_model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &fiber.Error{
			Code:    404,
			Message: "user not found",
		}
	}
	return user, nil
}
