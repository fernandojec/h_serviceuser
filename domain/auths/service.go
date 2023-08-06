package auths

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/fernandojec/h_serviceuser/config"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type irepo interface {
	GetAuthByEmail(email string) (auth *auths, err error)
	InsertRedis(c context.Context, key string, value interface{}, duration time.Duration) (err error)
	RemoveRedis(c context.Context, key string) (err error)
	GetRedis(c context.Context, key string) (value interface{}, err error)
}

type authsService struct {
	irepo irepo
}

func NewAuthsService(irepo irepo) *authsService {
	return &authsService{
		irepo: irepo,
	}
}

func (s *authsService) SignIn(ctx context.Context, data signInRequest) (dataResponse signInRequestResponse, err error) {
	dataUser, err := s.irepo.GetAuthByEmail(data.Email)
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

	accessToken, refreshToken, exp, err := s.generateTokenPair(ctx, dataUser)

	dataResponse = dataUser.ConvertToSignInResponse()
	dataResponse.RefreshToken = refreshToken
	dataResponse.AccessToken = accessToken
	dataResponse.ExpiresIn = exp

	return
}

func (s *authsService) generateTokenPair(ctx context.Context, auth *auths) (accessToken,
	refreshToken string,
	exp int64,
	err error,
) {
	var accessUID, refreshUID string
	if accessToken, accessUID, exp, err = s.createToken(auth, config.AppConfig.Jwt.ExpireAccessMinutes,
		ACCESS_SECRET); err != nil {
		return
	}

	if refreshToken, refreshUID, _, err = s.createToken(auth, config.AppConfig.Jwt.ExpireRefreshMinutes,
		REFRESH_SECRET); err != nil {
		return
	}

	cacheJSON, err := json.Marshal(CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	})
	s.irepo.InsertRedis(
		ctx,
		fmt.Sprintf("%s-token-%s", "H8-", auth.UserID),
		cacheJSON,
		(time.Minute * time.Duration(config.AppConfig.Jwt.AutoLogoffMinutes)),
	)
	return
}
func (s *authsService) createToken(user *auths, expireMinutes int, secret string) (token,
	uid string,
	exp int64,
	err error,
) {
	exp_time := time.Now().Add(time.Minute * time.Duration(expireMinutes))
	// exp = time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix()
	exp = exp_time.Unix()
	uid = uuid.New().String()

	claims := &MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   APPLICATION_NAME,
			Subject:  "",
			Audience: []string{},
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(constanta.LOGIN_EXPIRATION_DURATION)),
			ExpiresAt: jwt.NewNumericDate(exp_time),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        fmt.Sprint(user.UserID),
		},
		ID:        fmt.Sprint(user.UserID),
		UUID:      uid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))

	return
}
func (s *authsService) ParseToken(tokenString, secret string) (
	claims *MyClaims,
	err error,
) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
func (s *authsService) ValidateToken(c context.Context, claims *MyClaims, isRefresh bool) (

	user auths,
	err error,
) {
	var g errgroup.Group
	g.Go(func() error {
		cacheJSON, _ := s.irepo.GetRedis(c, fmt.Sprintf("%s-token-%s", "H8-", claims.ID))
		cachedTokens := new(CachedTokens)
		err = json.Unmarshal([]byte(cacheJSON.(string)), cachedTokens)

		var tokenUID string
		if isRefresh {
			tokenUID = cachedTokens.RefreshUID
		} else {
			tokenUID = cachedTokens.AccessUID
		}
		if err != nil || tokenUID != claims.UUID {
			return errors.New("token not found")
		}

		return nil
	})
	userChan := make([]auths, 1)
	if isRefresh {
		g.Go(func() error {
			var userT auths
			userTP, _ := s.irepo.GetAuthByEmail(claims.Email)
			if userTP == nil {
				return errors.New("user not found")
			}
			userT = *userTP
			userChan[0] = userT
			return nil
		})
	}
	err = g.Wait()
	user = userChan[0]
	return user, err
}
