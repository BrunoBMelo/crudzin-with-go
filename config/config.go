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
	} else {
		return &Config{
			StringConnection: "host=localhost port=5432 user=PG_USERGO password=PG_PWD2017 dbname=PG_DBGO sslmode=disable",
			ProviderName:     "postgres",
			PORT:             "3001",
		}
	}
	//panic("Erro when to load .env file")
}
