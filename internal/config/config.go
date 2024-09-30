package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ACCOMMODATION_SERVICE string
	DB_PASSWORD     string
	DB_HOST         string
	DB_PORT         int
	DB_USER         string
	DB_NAME         string
}
func Load() Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// ".env" faylining to'liq pathi
	if err := godotenv.Load(path + "/.env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "secret"))
	config.DB_PORT = cast.ToInt(Coalesce("DB_PORT", "secret"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "secret"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "secret"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "secret"))
	config.ACCOMMODATION_SERVICE = cast.ToString(Coalesce("ACCOMMODATION_SERVICE", "secret"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}