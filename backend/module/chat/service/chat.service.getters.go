package chat_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	chat_model "github.com/root9464/Ton-students/module/chat/model"
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

func (s *chatService) GetChatByID(ctx context.Context, chatID string) (*chat_model.Chat, error) {
	chat, err := s.repo.GetChatByID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
