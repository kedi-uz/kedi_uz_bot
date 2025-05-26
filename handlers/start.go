package handlers

import (
	"fmt"
	"kedi_uz_bot/buttons"
	"kedi_uz_bot/configs"
	"kedi_uz_bot/models"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"gorm.io/gorm"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	var user models.TelegramUser
	userData := user.GetUserData(ctx)

	var existingUser models.TelegramUser

	db := configs.DB 

	fmt.Println("------------------------------")
	fmt.Println(ctx.EffectiveUser.Id)
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println(userData.TelegramID)
	fmt.Println("------------------------------")
	
	if err := db.Where("telegram_id = ?", userData.TelegramID).First(&existingUser).Error; err == gorm.ErrRecordNotFound {
		db.Create(&userData)
	} else {
		db.Model(&userData).Where("telegram_id = ?", &userData.TelegramID).Updates(&userData)
	}

	_, err := b.SendMessage(ctx.EffectiveUser.Id, "choose district", &gotgbot.SendMessageOpts{
		ReplyMarkup: buttons.StartKeyboardMarkup(),
	})

	if err != nil {
		return fmt.Errorf("failed to send /about handler message: %w", err)
	}
	return handlers.NextConversationState(DISTRICT)
}


var validDistricts = map[string]bool{
	"Yunusobod":    true,
	"Chilonzor":    true,
	"Yakkasaroy":   true,
	"Shayxontohur": true,
	"Mirobod":    true,
	"Yashnobod":    true,
	"Olmazor":   true,
	"Mirzo Ulugʻbek": true,
	"Bektemir":    true,
	"Yangi Hayot":    true,
	"Sergeli":   true,
	"Yangi Toshkent": true,
	"Uchtepa": true,
}

// Create a matcher which only matches text which is ReplyMarkubButton
func NoCommands(msg *gotgbot.Message) bool {
	if msg == nil || msg.Text == "" {
		return false
	}
	
	// Only allow if message text matches one of the valid districts
	return validDistricts[msg.Text]
}

// // cancel cancels the conversation.
// func Cancel(b *gotgbot.Bot, ctx *ext.Context) error {
// 	_, err := ctx.EffectiveMessage.Reply(b, "Oh, goodbye!", &gotgbot.SendMessageOpts{
// 		ParseMode: "html",
// 	})
// 	if err != nil {
// 		return fmt.Errorf("failed to send cancel message: %w", err)
// 	}
// 	return handlers.EndConversation()
// }

// name gets the user's name.
func District(b *gotgbot.Bot, ctx *ext.Context) error {
	inputDistrict := ctx.EffectiveMessage.Text

	if !validDistricts[inputDistrict] {
		// If the number is not valid, try again!
		ctx.EffectiveMessage.Reply(b, "❌ Invalid district. Please select a district from the list below:", &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
		// We try the age handler again
		return handlers.NextConversationState(DISTRICT)
	}
	_, err := b.SendMessage(ctx.EffectiveUser.Id, fmt.Sprintf("Your district set to: %s!\n\nLost pets on this district will be notified to you", inputDistrict), &gotgbot.SendMessageOpts{
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