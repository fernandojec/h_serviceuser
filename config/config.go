package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Postgres Postgres
	App      App
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DBNAME"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}
type App struct {
	BaseUrl  string `env:"BASE_URL"`
	BasePort string `env:"BASE_PORT"`
}

func LoadConfig() (cfg Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		return
	}
	err = env.Parse(&cfg)
	return

}
