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


func (c *BotService) InlineQuery(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.InlineQuery.Query

	if query == "invite" {
		results := []gotgbot.InlineQueryResult{
			&gotgbot.InlineQueryResultArticle{
				Id:    "1", 
				Title: "Send Invite Link",
				InputMessageContent: &gotgbot.InputTextMessageContent{
					MessageText: "Click the button below to join!",
				},
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
						{
							{
								Text: "Join Now",                                     
								Url:  "https://t.me/ttonstudents_bot?start=12345678", 
							},
						},
					},
				},
			},
		}

		_, err := b.AnswerInlineQuery(ctx.InlineQuery.Id, results, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: 0, 
		})
		if err != nil {
			return fmt.Errorf("failed to answer inline query: %w", err)
		}
	}

	return nil
}
