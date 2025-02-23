package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type appConfig struct {
	AppPort string `env:"APP_PORT"`
	AppEnv  string `env:"APP_ENV"`
	AppUrl  string `env:"APP_URL"`

	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`

	JWTSecretKey string `env:"JWT_SECRET_KEY"`
}

var AppConfig appConfig

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading from system env")
	}

	if err := env.Parse(&AppConfig); err != nil {
		log.Fatal("Failed to parse environment variables: ", err)
	}
}
