package schedules

import (
	context "context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/config"
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB, redisc *redis.Client) {
	port := config.AppConfig.GrpcHost.ServiceSchedule
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ctx := context.Background()

		ctx = context.WithValue(ctx, loghelper.XTRACEID, fmt.Sprintf("%v", uuid.New()))
		loghelper.Errorf(ctx, "Error Connect To GRPC Server Schedules: %v", port)
	}
	service := NewScheduleClient(conn)
	handler := NewHandler(service)

	route := c.Group("schedules")

	route.Use(ifiber.ValidateJWT(dbx, redisc))
	route.Post("/", handler.CreateSchedule)
	route.Get("/FindScheduleAll", handler.FindScheduleAll)
	route.Get("/FindScheduleAvailable", handler.FindScheduleAvailable)
	route.Get("/FindScheduleBooked", handler.FindScheduleBooked)
	route.Put("/:schedule_id", handler.BookSlot)
}
