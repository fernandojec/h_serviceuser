package auths

import "github.com/golang-jwt/jwt/v4"

type signInRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type signInRequestResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
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
