package buttons

import "github.com/PaulSonOfLars/gotgbot/v2"

func StartKeyboardMarkup() gotgbot.ReplyKeyboardMarkup {
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard: [][]gotgbot.KeyboardButton{
			{
				{Text: "Yunusobod"},
				{Text: "Chilonzor"},
			},
			{
				{Text: "Yakkasaroy"},
				{Text: "Shayxontohur"},
			},
		},

		ResizeKeyboard: true,
		OneTimeKeyboard: true,
	}
}