package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramApiToken string
	psqlInfo string
}

func LoadConfig() *Config{
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not loaded, relying on system env vars")
	}
	token := os.Getenv("API_KEY")
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tashkent", host, user, password, dbname, port,)

	return &Config{
		TelegramApiToken: token,
		psqlInfo: psqlInfo,
	}
}