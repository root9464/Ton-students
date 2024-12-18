package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/ent"
)

func (s *userService) GetByID(ctx context.Context, id int64) (*ent.User, error) {
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
