package config

import (
	"context"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Postgres postgres
	App      app
	Jwt      jwt
	Redis    redisConfig
	GrpcHost grcpHost
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
	Ctx      context.Context
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

type grcpHost struct {
	ServiceParamedic   string `env:"SERVICE_PARAMEDIC"`
	ServicePatient     string `env:"SERVICE_PATIENT"`
	ServiceHealthcare  string `env:"SERVICE_HEALTHCARE"`
	ServiceAppointment string `env:"SERVICE_APPOINTMENT"`
	ServiceSchedule    string `env:"SERVICE_SCHEDULE"`
}

func LoadConfig() (cfg Config, err error) {
	_ = godotenv.Load(".env")
	err = env.Parse(&cfg)
	if err != nil {
		return
	}
	return

}
