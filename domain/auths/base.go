package auths

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB, redisc *redis.Client) {
	repo := NewAuthsRepo(dbx, redisc)
	service := NewAuthsService(repo)
	handler := NewAuthHandler(service)

	routeAuth := c.Group("auths")
	routeAuth.Post("/signin", handler.SignIn)
}
