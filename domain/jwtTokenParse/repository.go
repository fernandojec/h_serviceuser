package jwttokenparse

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repo struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewRepo(db *sqlx.DB, redisc *redis.Client) *Repo {
	return &Repo{db: db, redis: redisc}
}
func (r *Repo) GetAuthByEmail(email string) (data *AuthsJwt, err error) {
	qs := `select * from users where email=$1 and is_active = true`

	_, err = r.db.Prepare(qs)
	if err != nil {
		return
	}
	err = r.db.Get(data, qs, email)
	if err != nil {
		log.Errorf("Error db:%v", err)
	}
	return
}
func (r *Repo) InsertRedis(c context.Context, key string, value interface{}, duration time.Duration) (err error) {
	err = r.redis.Set(c, key, value, duration).Err()
	return
}
func (r *Repo) GetRedis(c context.Context, key string) (value interface{}, err error) {
	value, err = r.redis.Get(c, key).Result()
	return
}
