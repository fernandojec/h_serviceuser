package users

import (
	"context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/domain/notifications"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB) {
	repo := NewRepo(dbx)
	svc := NewService(repo)
	handler := NewHandler(svc)

	authApi := c.Group("/auth")
	authApi.Post("/sign-up", handler.CreateAuth)
	authApi.Get("/verify-email/:token", handler.ActivateAuth)
	authApi.Post("/sign-in", handler.SignInAuth)
}

func RouterInitUsers(c fiber.Router, dbx *sqlx.DB) {
	port := ":800"
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ctx := context.Background()

		ctx = context.WithValue(ctx, loghelper.XTRACEID, fmt.Sprintf("%v", uuid.New()))
		loghelper.Errorf(ctx, "Error Connect To GRPC Server Notification: %v", port)
	}
	repo_users := NewRepo_users(dbx)
	grpc_notification := notifications.NewMessagingServiceClient(conn)
	svc := NewServiceUsers(repo_users, grpc_notification)
	handler := NewHandler_Users(svc)

	usersAPI := c.Group("/users")
	usersAPI.Post("/saveregistrasi", handler.saveRegistrasi)
}
