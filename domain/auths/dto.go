package auths

type signInRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type signInRequestResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	ExpiresIn    int64  `json:"expires_in"`
}
type signOutRequest struct {
	AccessToken string `json:"access_token"`
}
