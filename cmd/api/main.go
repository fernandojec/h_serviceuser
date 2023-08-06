package main

import (
	"log"

	"github.com/fernandojec/h_serviceuser/config"
	"github.com/fernandojec/h_serviceuser/domain/auths"
	"github.com/fernandojec/h_serviceuser/domain/users"
	customvalidator "github.com/fernandojec/h_serviceuser/pkg/customValidator"
	"github.com/fernandojec/h_serviceuser/pkg/dbconnect"
	"github.com/fernandojec/h_serviceuser/pkg/redisconnect"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Cannot Load Config")
	}
	dbx, err := dbconnect.ConnectSqlx(dbconnect.DBConfig{
		Host:       cfg.Postgres.Host,
		Port:       cfg.Postgres.Port,
		Dbname:     cfg.Postgres.DbName,
		Dbuser:     cfg.Postgres.User,
		Dbpassword: cfg.Postgres.Password,
		Sslmode:    cfg.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("Cannot connect to DB:%v", err)
	}
	redisClient, err := redisconnect.ConnectRedis()
	if err != nil {
		log.Fatalf("Cannot connect to Redis:%v", err)
	}
	// _ = dbx
	app := fiber.New(
		fiber.Config{
			ErrorHandler: customvalidator.HttpErrorHandler,
		},
	)

	v1 := app.Group("v1")

	users.RouterInit(v1, dbx)
	auths.RouterInit(v1, dbx, redisClient)

	app.Listen(cfg.App.BasePort)
}
