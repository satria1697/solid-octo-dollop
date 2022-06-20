package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Jwt      JwtConfig
	Postgres PostgresConfig
}

type JwtConfig struct {
	Secret string
}

type PostgresConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

func GetConfig() Config {
	godotenv.Load()
	return Config{
		Jwt: JwtConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		Postgres: PostgresConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
			DbPort:     os.Getenv("DB_PORT"),
		},
	}
}
