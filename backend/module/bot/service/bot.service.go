package bot_service

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (c *BotService) InvateLink(b *gotgbot.Bot) (string, error) {
	invateLink, err := b.ExportChatInviteLink(c.config.ChatBotId, nil)

	if err != nil {
		c.logger.Error("Failed to execute InvateLink command: " + err.Error())
		return "", err
	}

	return invateLink, nil
}

// func (c *BotService) HandleInlineQuery(bot *gotgbot.Bot, query *gotgbot.InlineQuery) error {
// 	groupChatID := int64(-1001234567890)

// 	inviteLink, err := bot.ExportChatInviteLink(groupChatID, nil)
// 	if err != nil {
// 		log.Printf("Ошибка генерации invite-ссылки: %v", err)
// 		return nil
// 	}

// 	results := []gotgbot.InlineQueryResult{
// 		&gotgbot.InlineQueryResultArticle{
// 			Id:    "invite_link",
// 			Title: "Ссылка приглашения в группу",
// 			InputMessageContent: gotgbot.InputTextMessageContent{
// 				MessageText: fmt.Sprintf("Вот ваша ссылка приглашения: %s", inviteLink),
// 			},
// 			Description: "Нажмите, чтобы получить ссылку приглашения в группу.",
// 		},
// 	}

// 	_, err = query.Answer(bot, results, &gotgbot.AnswerInlineQueryOpts{
// 		CacheTime: 0,
// 	})
// 	if err != nil {
// 		log.Printf("Ошибка отправки inline-ответа: %v", err)
// 	}

// 	return nil
// }

func (c *BotService) InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	// Получаем текст запроса пользователя
	query := ctx.InlineQuery.Query

	// Проверяем, что запрос содержит "invite"
	if query == "invite" {
		// Создаем inline-результат с кнопкой
		results := []gotgbot.InlineQueryResult{
			&gotgbot.InlineQueryResultArticle{
				Id:    "1", // уникальный ID результата
				Title: "Send Invite Link",
				InputMessageContent: &gotgbot.InputTextMessageContent{
					MessageText: "Click the button below to join!",
				},
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
						{
							{
								Text: "Join Now",                                     // Текст кнопки
								Url:  "https://t.me/ttonstudents_bot?start=12345678", // Ссылка на приглашение
							},
						},
					},
				},
			},
		}

		// Отправляем inline-результаты
		_, err := b.AnswerInlineQuery(ctx.InlineQuery.Id, results, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: 0, // Обнуляем кэш, чтобы изменения были видны сразу
		})
		if err != nil {
			return fmt.Errorf("failed to answer inline query: %w", err)
		}
	}

	return nil
}
