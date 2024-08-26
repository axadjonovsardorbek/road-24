package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	MOBILE_PORT string
	MOBILE_HOST string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	REDIS_HOST string
	REDIS_PORT string

	KAFKA_HOST string
	KAFKA_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.MOBILE_PORT = cast.ToString(coalesce("AUTH_PORT", ":50908"))
	config.MOBILE_HOST = cast.ToString(coalesce("AUTH_HOST", "mobile-road"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1111"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "delivery"))

	config.REDIS_HOST = cast.ToString(coalesce("REDIS_HOST", "localhost"))
	config.REDIS_PORT = cast.ToString(coalesce("REDIS_PORT", ":6379"))

	config.KAFKA_HOST = cast.ToString(coalesce("KAFKA_HOST", "localhost"))
	config.KAFKA_PORT = cast.ToString(coalesce("KAFKA_PORT", ":9092"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
