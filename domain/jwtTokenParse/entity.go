package jwttokenparse

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthsJwt struct {
	UserID     string    `db:"user_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	IsActive   bool      `db:"is_active"`
	UserCreate bool      `db:"user_create"`
	CreateAt   time.Time `db:"create_at"`
}

type MyClaims struct {
	jwt.RegisteredClaims
	ID        string `json:"id"`
	UUID      string `json:"uuid"`
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}
