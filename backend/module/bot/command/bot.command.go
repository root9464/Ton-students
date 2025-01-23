package bot_command

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (cm *botCommand) StartMessage(bot *gotgbot.Bot, ctx *ext.Context) error {
	username := ctx.EffectiveUser.FirstName
	if username == "" {
		username = "there"
	}

	messageText := fmt.Sprintf("Hello, %s! 👋\nWelcome to the bot. How can I assist you today?", username)

	_, err := ctx.EffectiveMessage.Reply(bot, messageText, nil)
	if err != nil {
		cm.logger.Error("Failed to send start message: " + err.Error())
		return err
	}

	return nil
}
