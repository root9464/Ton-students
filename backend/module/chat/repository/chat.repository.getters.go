package chat_repository

import (
	"context"
	"strings"

	chat_model "github.com/root9464/Ton-students/module/chat/model"
)

func (r *chatRepository) GetChatIDBetweenUsers(ctx context.Context, userIDs []int64) (*string, error) {
	var chatID string

	placeholders := make([]string, len(userIDs))
	for i := range userIDs {
		placeholders[i] = "?"
	}

	interfaceIDs := make([]interface{}, len(userIDs))
	for i, id := range userIDs {
		interfaceIDs[i] = id
	}

	err := r.db.WithContext(ctx).Table("chat_users cu1").
		Select("cu1.chat_id").
		Joins("JOIN chat_users cu2 ON cu1.chat_id = cu2.chat_id").
		Where("cu1.user_id IN ("+strings.Join(placeholders, ",")+")", interfaceIDs...).
		Group("cu1.chat_id").
		Having("COUNT(DISTINCT cu1.user_id) = ?", len(userIDs)).
		Limit(1). // Ограничиваем результат до одного чата
		Scan(&chatID).Error

	if err != nil {
		return nil, err
	}

	if chatID == "" {
		return nil, nil // Если чат не найден, возвращаем nil
	}

	return &chatID, nil
}

func (r *chatRepository) GetChatByID(ctx context.Context, chatID string) (*chat_model.Chat, error) {
	chat := new(chat_model.Chat)
	if err := r.db.WithContext(ctx).Where("id = ?", chatID).First(chat).Error; err != nil {
		return nil, err
	}
	return chat, nil
}
