package auths

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type authsRepo struct {
	db *sqlx.DB
	redis *redis.Client
}

func NewAuthsRepo(db *sqlx.DB) *authsRepo {
	return &authsRepo{db: db}
}

func (r *authsRepo) GetAuthByEmail(email string) (data *auths, err error) {
	qs := `select * from users where email=$1`

	_, err = r.db.Prepare(qs)
	if err != nil {
		return
	}
	err = r.db.Get(data, qs, email)
	return
}

func (r *authsRepo) InsertRedis(c context.Context, key string, value interface{}, duration time.Duration){
	
}