package models

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TelegramUser struct {
	ID uint `gorm:"primaryKey"`
	TelegramID int64 `gorm:"unique"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	UserName string
	FirstName string
	LastName string
	LanguageCode string
	District string
}

func (u *TelegramUser) GetUserData(ctx *ext.Context) TelegramUser {
	user := ctx.EffectiveUser
	u.TelegramID = user.Id
	u.UserName = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.LanguageCode = user.LanguageCode

	return *u
}