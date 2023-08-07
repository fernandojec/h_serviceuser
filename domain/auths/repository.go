package auths

import (
	"context"
	"time"

	"github.com/fernandojec/h_serviceuser/pkg/loghelper"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type authsRepo struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewAuthsRepo(db *sqlx.DB, redis *redis.Client) *authsRepo {
	return &authsRepo{db: db, redis: redis}
}

func (r *authsRepo) GetAuthByEmail(ctx context.Context, email string) (data *auths, err error) {
	qs := `select * from users where email=$1 and is_active = true`

	_, err = r.db.Prepare(qs)
	if err != nil {
		return
	}
	data = &auths{}
	err = r.db.Get(data, qs, email)
	if err != nil {
		data = nil
		loghelper.Errorf(ctx, "Error Get Data From DB :%v", err)
	}
	return
}

func (r *authsRepo) InsertRedis(c context.Context, key string, value interface{}, duration time.Duration) (err error) {
	err = r.redis.Set(c, key, value, duration).Err()
	return
}

func (r *authsRepo) RemoveRedis(c context.Context, key string) (err error) {
	err = r.redis.Del(c, key).Err()
	return
}
func (r *authsRepo) GetRedis(c context.Context, key string) (value interface{}, err error) {
	value, err = r.redis.Get(c, key).Result()
	return
}
