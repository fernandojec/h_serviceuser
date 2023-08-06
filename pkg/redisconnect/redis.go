package redisconnect

import (
	"context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (redisClient *redis.Client, err error) {
	ctx := context.Background()
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Host),
			Password: config.AppConfig.Redis.Password,
		},
	)
	err = redisClient.Ping(ctx).Err()
	if err != nil {
		return
	}
	return
}