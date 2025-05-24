package buttons

import "github.com/PaulSonOfLars/gotgbot/v2"

func StartKeyboardMarkup() gotgbot.ReplyKeyboardMarkup {
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard: [][]gotgbot.KeyboardButton{
			{
				{Text: "Yunusobod"},
				{Text: "Olmazor"},
				{Text: "Shayxontohur"},
			},
			{
				{Text: "Uchtepa"},
				{Text: "Yakkasaroy"},
				{Text: "Chilonzor"},
			},
			{
				{Text: "Mirobod"},
				{Text: "Yashnobod"},
				{Text: "Mirzo Ulugʻbek"},
			},
			{
				{Text: "Sergeli"},
				{Text: "Yangi Hayot"},
				{Text: "Bektemir"},
			},
			{
				{Text: "Yangi Toshkent"},
			},
		},

		ResizeKeyboard: true,
		OneTimeKeyboard: true,
	}
}