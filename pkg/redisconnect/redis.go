package redisconnect

import (
	"context"
	"fmt"

	"github.com/fernandojec/h_serviceuser/config"
	nrredis "github.com/newrelic/go-agent/v3/integrations/nrredis-v9"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (redisClient *redis.Client, err error) {
	ctx := context.Background()
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port),
		Password: config.AppConfig.Redis.Password,
		DB:       0,
	}
	redisClient = redis.NewClient(
		opts,
	)
	redisClient.AddHook(nrredis.NewHook(opts))

	err = redisClient.Ping(ctx).Err()
	if err != nil {
		return
	}
	return
}
