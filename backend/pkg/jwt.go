package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
	secret          = os.Getenv("JWT_SECRET")
)

type Pair struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}

func NewPair(userID string) (*Pair, error) {
	jti := uuid.NewString()

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
		ID:        jti,
	})

	accessStr, err := access.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenTTL)),
		ID:        jti,
	})

	refreshStr, err := refresh.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &Pair{AccessToken: accessStr, RefreshToken: refreshStr, UserID: userID}, nil
}

func Parse(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*jwt.RegisteredClaims), nil
}
