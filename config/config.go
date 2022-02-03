package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StringConnection string
	ProviderName     string
	PORT             string
}

func Load() *Config {
	if err := godotenv.Load(); err == nil {
		return &Config{
			StringConnection: os.Getenv("PG_STRT_CONNECTION"),
			ProviderName:     os.Getenv("PROVIDERNAME"),
			PORT:             os.Getenv("APP_PORT"),
		}
	}
	panic("Erro when to load .env file")
}
