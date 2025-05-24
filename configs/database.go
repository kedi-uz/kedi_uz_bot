package configs

import (
	"kedi_uz_bot/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	psqlInfo := LoadConfig().psqlInfo

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not loaded, relying on system env vars")
	}

	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	
	if err != nil {
		log.Fatal("failed to start Databse: " + err.Error())
	}

	err = DB.AutoMigrate(&models.TelegramUser{})
	if err != nil {
		log.Fatal("failed to start automigration: " + err.Error())
	}
	
	return DB
}

