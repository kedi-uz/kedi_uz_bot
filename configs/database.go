package configs

import (
	"kedi_uz_bot/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	
	dsn := "host=localhost user=postgres password=postgres dbname=kedi_uz_bot port=5432 sslmode=disable TimeZone=Asia/Tashkent"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("failed to start Databse: " + err.Error())
	}

	err = DB.AutoMigrate(&models.TelegramUser{})
	if err != nil {
		log.Fatal("failed to start automigration: " + err.Error())
	}
	
	return DB
}

