package chat_repository

import (
	"context"

	chat_model "github.com/root9464/Ton-students/module/chat/model"
	common_model "github.com/root9464/Ton-students/module/model/common"
)

func (r *chatRepository) CreateChat(ctx context.Context, chat *chat_model.Chat) error {
	r.logger.Info("Creating chat...")
	if err := r.db.WithContext(ctx).Create(&chat).Error; err != nil {
		r.logger.Errorf("Error creating chat: %v", err)
		return err
	}
	r.logger.Info("Chat create successfully")
	return nil
}

func (r *chatRepository) CreateChatMembers(ctx context.Context, chatMembers []common_model.ChatUser) error {
	r.logger.Info("Creating chat members...")
	if err := r.db.WithContext(ctx).Create(&chatMembers).Error; err != nil {
		r.logger.Errorf("Error creating chat members: %v", err)
		return err
	}
	r.logger.Info("Chat members create successfully")
	return nil
}
