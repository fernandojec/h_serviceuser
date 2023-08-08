package ifiber

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fernandojec/h_serviceuser/config"
	jwttokenparse "github.com/fernandojec/h_serviceuser/domain/jwtTokenParse"
	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func InsertContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		xTraceID := c.Get("x-trace-id")
		ctx := c.UserContext()
		ctx = context.WithValue(ctx, loghelper.XTRACEID, xTraceID)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func ValidateJWT(dbx *sqlx.DB, redisc *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("x-token")
		if token == "" {
			return ErrorResponse(c, fiber.StatusUnauthorized, fiber.Map{
				"error":   "unauthenticated",
				"message": "x-token is not provide",
			})
		}
		repo := jwttokenparse.NewRepo(dbx, redisc)
		tokenService := jwttokenparse.NewService(repo)
		claims, err := tokenService.ParseToken(token, jwttokenparse.ACCESS_SECRET)
		if err != nil {
			return ErrorResponse(c, fiber.StatusUnauthorized, fiber.Map{
				"error":        "unauthenticated",
				"message":      err.Error(),
				"addt-message": "Error Parse Token",
			})
		}
		_, err = tokenService.ValidateToken(c.UserContext(), claims, false)
		if err != nil {
			return ErrorResponse(c, fiber.StatusUnauthorized, fiber.Map{
				"error":        "unauthenticated",
				"message":      err.Error(),
				"addt-message": "Error Validate Token",
			})
		}

		go func() {
			redisc.Expire(c.Context(), fmt.Sprintf("%s-token-%s", "H8", claims.ID),
				time.Duration(60*config.AppConfig.Jwt.AutoLogoffMinutes))
		}()
		id, _ := strconv.ParseUint(claims.ID, 10, 64)
		c.Locals("UserID", uint(id))
		ctx := c.UserContext()
		ctx = context.WithValue(ctx, USERID, id)
		c.SetUserContext(ctx)
		return c.Next()
	}
}
