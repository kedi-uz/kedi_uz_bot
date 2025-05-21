package handlers

import (
	"fmt"
	"kedi_uz_bot/configs"
	"kedi_uz_bot/models"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gorm.io/gorm"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	var user models.TelegramUser
	userData := user.GetUserData(ctx)

	db := configs.DB 

	if err := db.First(&userData).Where("telegram_id = ?", userData.TelegramID).Error; err == gorm.ErrRecordNotFound {
		db.Create(&userData)
	} else {
		db.Model(&userData).Where("telegram_id = ?", &userData.TelegramID).Updates(&userData)
	}

	_, err := ctx.EffectiveMessage.Reply(b, "/start", &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})

	if err != nil {
		return fmt.Errorf("failed to send /about handler message: %w", err)
	}
	return nil
}

