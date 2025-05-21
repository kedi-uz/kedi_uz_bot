package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramApiToken string
}

func LoadConfig() *Config{
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not loaded, relying on system env vars")
	}
	
	return &Config{
		TelegramApiToken: os.Getenv("API_KEY"),
	}
}

// Load single config variable but it repeat the godotenv load when every time called
// func LoadConfig(key string) string {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Println("Warning: .env file not loaded, relying on system env vars")
// 	}

// 	return os.Getenv(key)
// }