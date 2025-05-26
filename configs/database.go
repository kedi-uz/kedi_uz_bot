package configs

import (
	"kedi_uz_bot/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error

	psqlInfo := LoadConfig().psqlInfo

	// Retry loop to wait for DB to be ready
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("🔁 Waiting for database... attempt %d/10: %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}
	
	if err != nil {
		log.Fatal("failed to start Databse: " + err.Error())
	}

	log.Println("✅ Connected to the database!")
	
	err = DB.AutoMigrate(&models.TelegramUser{})
	if err != nil {
		log.Fatal("failed to start automigration: " + err.Error())
	}
	
	return DB
}

