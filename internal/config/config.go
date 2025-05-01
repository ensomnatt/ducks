package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	PostgresUser string
	PostgresPassword string
}

func GetConfig() Config {
	godotenv.Load(".env")

	return Config{
		Env: os.Getenv("ENV"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
}
