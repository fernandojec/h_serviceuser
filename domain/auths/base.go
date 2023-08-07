package auths

import (
	jwttokenparse "github.com/fernandojec/h_serviceuser/domain/jwtTokenParse"
	"github.com/fernandojec/h_serviceuser/infra/ifiber"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func RouterInit(c fiber.Router, dbx *sqlx.DB, redisc *redis.Client) {

	jwtRepo := jwttokenparse.NewRepo(dbx, redisc)
	jwtService := jwttokenparse.NewService(jwtRepo)

	repo := NewAuthsRepo(dbx, redisc)
	service := NewAuthsService(repo, *jwtService)
	handler := NewAuthHandler(service)

	routeAuth := c.Group("auths")
	routeAuth.Post("/signin", handler.SignIn)
	routeAuth.Post("/signout", ifiber.ValidateJWT(dbx, redisc), handler.SignOut)
}
