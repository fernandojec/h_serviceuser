package auths

import "time"

type auths struct {
	UserID     string    `db:"user_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	IsActive   bool      `db:"is_active"`
	UserCreate bool      `db:"user_create"`
	CreateAt   time.Time `db:"create_at"`
}

func (a auths) ConvertToSignInResponse() signInRequestResponse {
	return signInRequestResponse{
		UserID:    a.UserID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Email:     a.Email,
	}
}
