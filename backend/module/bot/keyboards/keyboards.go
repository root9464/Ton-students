package bot_keyboards

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

var Keyboard = gotgbot.ReplyKeyboardMarkup{
	Keyboard: [][]gotgbot.KeyboardButton{
		{
			{Text: "/start" },
			{Text: "/news"},
			{Text: "/support"},
			
		},		
	},
	ResizeKeyboard: true, // Подгоняем под экран
}
