package hospital

import (
	context "context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/config"
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB, redisc *redis.Client, appNewRelic *newrelic.Application) {
	port := config.AppConfig.GrpcHost.ServiceHealthcare
	conn, err := grpc.Dial(
		port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(nrgrpc.StreamClientInterceptor),
	)
	if err != nil {
		ctx := context.Background()

		ctx = context.WithValue(ctx, loghelper.XTRACEID, fmt.Sprintf("%v", uuid.New()))
		loghelper.Errorf(ctx, "Error Connect To GRPC Server Hospital: %v", port)
	}
	service := NewHospitalServiceClient(conn)
	handler := NewHandler(service, appNewRelic)

	route := c.Group("hospitals")

	route.Use(ifiber.ValidateJWT(dbx, redisc))
	route.Get("/", handler.List)
	route.Get("/:id", handler.Get)
	route.Post("/", handler.Create)
}
