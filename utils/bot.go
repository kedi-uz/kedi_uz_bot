package utils

import (
	"fmt"
	"html"
	"kedi_uz_bot/configs"
	"kedi_uz_bot/models"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var validDistricts = map[string]bool{
	"Yunusobod":    true,
	"Chilonzor":    true,
	"Yakkasaroy":   true,
	"Shayxontohur": true,
}

// Create a matcher which only matches text which is ReplyMarkubButton
func NoCommands(msg *gotgbot.Message) bool {
	if msg == nil || msg.Text == "" {
		return false
	}
	// Only allow if message text matches one of the valid districts
	return validDistricts[msg.Text]
}

// cancel cancels the conversation.
func Cancel(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "Oh, goodbye!", &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send cancel message: %w", err)
	}
	return handlers.EndConversation()
}

// name gets the user's name.
func District(b *gotgbot.Bot, ctx *ext.Context) error {
	inputDistrict := ctx.EffectiveMessage.Text
	_, err := b.SendMessage(ctx.EffectiveUser.Id, fmt.Sprintf("Your district set to: %s!\n\nLost pets on this district will be notified to you", html.EscapeString(inputDistrict)), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to send name message: %w", err)
	}

	var user models.TelegramUser
	userData := user.GetUserData(ctx)

	db := configs.DB 
	db.Model(&user).Where("telegram_id = ?", userData.TelegramID).Update("district", inputDistrict)

	return handlers.EndConversation()
}