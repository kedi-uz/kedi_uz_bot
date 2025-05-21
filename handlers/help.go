package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)


func help(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, getHelpText(), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})

	if err != nil {
		return fmt.Errorf("failed to send /help handler message: %w", err)
		
	}
	return nil
}

func getHelpText() string {
	return `Mavjud komandalar ro'yxati:

/help - ushbu xabarni qayta ko'rsatish
/about - ushbu bot haqida ma'lumot
/stats - kedi.uz haqida statistika ma'lumotlari
`
}