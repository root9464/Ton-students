package chat_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (s *chatService) GetChatIDBetweenUsers(ctx context.Context, userIDs []int64) (*string, error) {
	if len(userIDs) < 2 {
		return nil, &fiber.Error{
			Code:    400,
			Message: "At least two users are required",
		}
	}
	chatID, err := s.repo.GetChatIDBetweenUsers(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return chatID, nil
}
