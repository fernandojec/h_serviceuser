package auths

import (
	"context"
	"database/sql"
	"errors"
	"time"

	jwttokenparse "github.com/fernandojec/h_serviceuser/domain/jwtTokenParse"
	"golang.org/x/crypto/bcrypt"
)

type irepo interface {
	GetAuthByEmail(email string) (auth *auths, err error)
	InsertRedis(c context.Context, key string, value interface{}, duration time.Duration) (err error)
	RemoveRedis(c context.Context, key string) (err error)
	GetRedis(c context.Context, key string) (value interface{}, err error)
}

type authsService struct {
	irepo      irepo
	jwtService jwttokenparse.Service
}

func NewAuthsService(irepo irepo, jwtService jwttokenparse.Service) *authsService {
	return &authsService{
		irepo:      irepo,
		jwtService: jwtService,
	}
}

func (s *authsService) SignIn(ctx context.Context, data signInRequest) (dataResponse signInRequestResponse, err error) {
	dataUser, err := s.irepo.GetAuthByEmail(data.Email)
	if err != nil && err == sql.ErrNoRows {
		err = errors.New("user not found")
		return
	}
	if err != nil {
		return
	}
	if dataUser == nil {
		err = errors.New("user not found")
		return
	}
	if !dataUser.IsActive {
		err = errors.New("user not active")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password)); err != nil {
		err = errors.New("invalid password")
		return
	}

	dataAuthJWT := dataUser.ConvertToAuthJWT()

	accessToken, refreshToken, exp, err := s.jwtService.GenerateTokenPair(ctx, &dataAuthJWT)

	dataResponse = dataUser.ConvertToSignInResponse()
	dataResponse.RefreshToken = refreshToken
	dataResponse.AccessToken = accessToken
	dataResponse.ExpiresIn = exp

	return
}

func (s *authsService) SignOut(ctx context.Context, data signOutRequest) (err error) {

	claims, err := s.jwtService.ParseToken(data.AccessToken, jwttokenparse.ACCESS_SECRET)
	if err != nil {
		return err
	}
	err = s.irepo.RemoveRedis(ctx, claims.ID)
	return
}
