package chat_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	chat_dto "github.com/root9464/Ton-students/module/chat/dto"
	chat_model "github.com/root9464/Ton-students/module/chat/model"
	common_model "github.com/root9464/Ton-students/module/model/common"
)

func (s *chatService) CreateChat(ctx context.Context, dto *chat_dto.CreateChatType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	chatModel := new(chat_model.Chat)

	// hash := sha256.New()
	// hash.Write([]byte(strconv.FormatInt(dto.Users[0]+dto.Users[1], 10)))
	// key := string(hash.Sum(nil))

	key := "qwe"

	chatModel.Key = key
	if err := s.repo.CreateChat(ctx, chatModel); err != nil {
		s.logger.Errorf("create chat error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	chatMembers := []common_model.ChatUser{
		{
			UserID: dto.Users[0],
			ChatID: chatModel.ID,
		},
		{
			UserID: dto.Users[1],
			ChatID: chatModel.ID,
		},
	}

	if err := s.repo.CreateChatMembers(ctx, chatMembers); err != nil {
		s.logger.Errorf("create chat members error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *chatService) CreateOrLoadChat(ctx context.Context, dto *chat_dto.CreateOrLoad) error {
	s.logger.Infof("dto: %v", dto)
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	service, err := s.serviceService.GetServiceById(ctx, dto.ServiceID)
	if err != nil {
		return err
	}

	userIDs := []int64{service.UserID, dto.UserID}
	chatID, err := s.repo.GetChatIDBetweenUsers(ctx, userIDs)
	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if chatID == nil {
		return s.CreateChat(ctx, &chat_dto.CreateChatType{
			Users: userIDs,
		})
	}

	return nil
}
