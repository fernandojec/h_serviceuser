package main

import (
	"context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/config"
	"github.com/fernandojec/h_serviceuser/domain/appointment"
	"github.com/fernandojec/h_serviceuser/domain/auths"
	"github.com/fernandojec/h_serviceuser/domain/hospital"
	"github.com/fernandojec/h_serviceuser/domain/paramedics"
	"github.com/fernandojec/h_serviceuser/domain/patient"
	"github.com/fernandojec/h_serviceuser/domain/schedules"
	"github.com/fernandojec/h_serviceuser/domain/users"
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	customvalidator "github.com/fernandojec/h_serviceuser/pkg/customValidator"
	"github.com/fernandojec/h_serviceuser/pkg/dbconnect"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/fernandojec/h_serviceuser/pkg/redisconnect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, loghelper.XTRACEID, fmt.Sprintf("Init%v", uuid.New()))
	cfg, err := config.LoadConfig()
	cfg.App.Ctx = ctx
	config.AppConfig = cfg

	if err != nil {
		loghelper.Infof(ctx, "Cannot load config: %v", err)
		// panic("Cannot Load Config")
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
		loghelper.Fatalf(ctx, "Cannot connect to DB:%v", err)
	}
	redisClient, err := redisconnect.ConnectRedis()
	if err != nil {
		loghelper.Fatalf(ctx, "Cannot connect to Redis:%v", err)
	}
	// _ = dbx
	app := fiber.New(
		fiber.Config{
			ErrorHandler: customvalidator.HttpErrorHandler,
		},
	)
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	v1 := app.Group("v1")
	v1.Use(ifiber.InsertContext())

	users.RouterInitUsers(v1, dbx)
	auths.RouterInit(v1, dbx, redisClient)
	paramedics.RouterInit(v1, dbx, redisClient)
	patient.RouterInit(v1, dbx, redisClient)
	appointment.RouterInit(v1, dbx, redisClient)
	schedules.RouterInit(v1, dbx, redisClient)
	hospital.RouterInit(v1, dbx, redisClient)
	//kkoo
	fmt.Println("test ricky")

	app.Listen(cfg.App.BasePort)
}
