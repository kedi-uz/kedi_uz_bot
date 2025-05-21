package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func about(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, getAboutText(), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})

	if err != nil {
		return fmt.Errorf("failed to send /about handler message: %w", err)
	}
	return nil
}

func getAboutText() string {
	return `🌆Toshkent mushuklar shahri

🇹🇷Istanbul - mushuklar shahri. Butun dunyodan sayyohlar mushuklar shahrini ko'rish uchun har yili Istanbulga kelishadi va u yerda yashaydigan odamlar haqida yaxshi taassurot qoldiriib ketishadi.

🐈Chunki mahaliy aholi mushuklarga  juda mehribon, ularga ovqat berishadi, kafe va restoranga mushuklar kirishsa xaydab yuborishmaydi.

🇺🇿O'zbekistonni ham shundan hayvonlar va insonlar birga ahil yashaydigan davlat qilsa bo'ladi agar ularga mehribonlik qilsak.

🧑‍💻Website: kedi.uz
Telegram Bot: @kedi_uz_bot
📢Telegram kanalga obuna bo'ling: t.me/kedi_uz
`
}
