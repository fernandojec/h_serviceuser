package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Postgres postgres
	App      app
	Jwt      jwt
	Redis    redisConfig
}

type postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DBNAME"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}
type app struct {
	BaseUrl  string `env:"BASE_URL"`
	BasePort string `env:"BASE_PORT"`
}

type jwt struct {
	ExpireAccessMinutes  int `env:"EXPIRE_ACCESS_MIN"`
	ExpireRefreshMinutes int `env:"EXPIRE_REFRESH_MIN"`
	AutoLogoffMinutes    int `env:"AUTO_LOGOFF_MIN"`
}

type redisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

func LoadConfig() (cfg Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}
	err = env.Parse(&cfg)
	return

}
